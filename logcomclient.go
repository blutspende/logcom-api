package logcomapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom/log"
	"github.com/google/uuid"
)

var (
	configuration  LogComConfiguration
	instance       *APIClient
	internalLogger zerolog.Logger
	once           sync.Once
)

type LogComConfiguration struct {
	ServiceName    string
	LogComURL      string
	HeaderProvider HeaderProviderFunc
}

type HeaderProviderFunc func(ctx context.Context) http.Header

func Init(config LogComConfiguration, logger *zerolog.Logger) {
	if config.LogComURL == "" {
		logger.Error().Msg("LogCom URL is missing thus functionalities are not available")
		return
	}

	once.Do(func() {
		configuration = config

		logComAPIConfig := NewConfiguration()
		logComAPIConfig.BasePath = configuration.LogComURL + "/api"
		logComAPIConfig.Host = configuration.LogComURL

		internalLogger = logger.Sample(zerolog.LevelSampler{
			ErrorSampler: &zerolog.BurstSampler{
				Burst:       1,
				Period:      5 * time.Second,
				NextSampler: &zerolog.BasicSampler{N: 20},
			},
		})

		parsedUrl, err := url.Parse(configuration.LogComURL)
		if err != nil {
			internalLogger.Error().Err(err).
				Str("url", configuration.LogComURL).
				Msgf("Failed to get LogCom URL scheme, falling back to default (%s)", "https")

			parsedUrl = &url.URL{
				Scheme: "https",
			}
		}

		logComAPIConfig.Scheme = parsedUrl.Scheme

		instance = NewAPIClient(logComAPIConfig)
		*logger = logger.Hook(logComHook{internalLogger: internalLogger})
	})
}

func IsEnabled() bool {
	return instance != nil
}

func SendConsoleLog(ctx context.Context, logLevel zerolog.Level, message string) error {
	return SendConsoleLogWithModel(ctx, CreateConsoleLogRequestDto{
		Level:   int32(logLevel),
		Message: message,
		Service: configuration.ServiceName,
	})
}

func SendConsoleLogWithModel(ctx context.Context, model CreateConsoleLogRequestDto) error {
	// Not LogCom intended log message
	if ctx == context.TODO() || ctx == context.Background() {
		return nil
	}

	if !IsEnabled() {
		internalLogger.Debug().Msg("LogCom is disabled")
		return nil
	}

	prepareConsoleLogRequestDTO(&model)

	headers := configuration.HeaderProvider(ctx)
	result, err := instance.ConsoleLogApi.CreateConsoleLogV1Int(ctx, model, requestConfigurer(ctx, headers))
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func SendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{}) error {
	if newValue == nil {
		newValue = ""
	}
	return SendAuditLog(ctx, CreateAuditLogRequestDto{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     subject,
		SubjectName: subjectName,
	})
}

func SendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if isPrimitiveType(oldValue) {
		return SendAuditLog(ctx, CreateAuditLogRequestDto{
			Category:    "MODIFICATION",
			NewValue:    fmt.Sprintf("%v", newValue),
			OldValue:    fmt.Sprintf("%v", oldValue),
			Subject:     subject,
			SubjectName: subjectName,
		})
	}

	changes, err := GetModelChanges(ctx, oldValue, newValue, ignoredProperties...)
	if err != nil {
		log.Error().MsgContext(ctx, "Failed to send audit log")
		return err
	}

	if len(changes) < 1 {
		log.Debug().MsgContext(ctx, "No changes")
		return nil
	}

	return SendAuditLogWithModificationModelChanges(ctx, subject, subjectName, changes)
}

func SendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange) error {
	dto := CreateAuditLogRequestDto{
		Category:    "MODIFICATION",
		Subject:     subject,
		SubjectName: subjectName,
	}

	if changesCount := len(changes); changesCount > 1 {
		changesDTO := make([]NewAuditLogChangeDto, changesCount)

		for i, change := range changes {
			changesDTO[i] = NewAuditLogChangeDto{
				Category:            dto.Category,
				NewValue:            stringify(change.NewValue),
				OldValue:            stringify(change.OldValue),
				Subject:             dto.Subject,
				SubjectName:         dto.SubjectName,
				SubjectPropertyName: change.PropertyName,
			}
		}

		dto.GroupedChanges = changesDTO
	} else if changesCount > 0 {
		dto.NewValue = stringify(changes[0].NewValue)
		dto.OldValue = stringify(changes[0].OldValue)
		dto.SubjectPropertyName = changes[0].PropertyName
	}

	return SendAuditLog(ctx, dto)
}

func SendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{}) error {
	if oldValue == nil {
		oldValue = ""
	}
	return SendAuditLog(ctx, CreateAuditLogRequestDto{
		Category:    "DELETION",
		OldValue:    stringify(oldValue),
		Subject:     subject,
		SubjectName: subjectName,
	})
}

func SendAuditLog(ctx context.Context, model CreateAuditLogRequestDto) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}

	if !IsEnabled() {
		internalLogger.Info().Msg("LogCom is disabled")
		return nil
	}

	prepareAuditLogRequestDTO(&model)

	headers := configuration.HeaderProvider(ctx)
	result, err := instance.AuditLogApi.CreateAuditLogV1Int(ctx, model, requestConfigurer(ctx, headers))
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func SendAuditLogGroup(ctx context.Context, auditLogCollector *AuditLogCollector) error {
	return SendAuditLog(ctx, auditLogCollector.get())
}

func isHTTPStatusSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}

func requestConfigurer(ctx context.Context, headers http.Header) func(*http.Request) {
	return func(request *http.Request) {
		requestID := headers.Get("X-Request-ID")
		if requestID == "" {
			if requestIDAsUUID, ok := ctx.Value("RequestID").(uuid.UUID); ok {
				requestID = requestIDAsUUID.String()
			}
		}

		if requestID == "" {
			request.Header.Add("X-Request-ID", requestID)
		}

		authorization := headers.Get("Authorization")
		if authorization != "" {
			request.Header.Add("Authorization", authorization)
		}
	}
}

func stringify(value interface{}) string {
	if isPrimitiveType(value) {
		return fmt.Sprintf("%v", value)
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%+v", value)
	}

	return string(jsonData)
}

func isPrimitiveType(value interface{}) bool {
	valueKind := reflect.TypeOf(value).Kind()
	return ((valueKind < reflect.Array) && (valueKind > 0) && (valueKind != reflect.Uintptr)) || (valueKind == reflect.String)
}
