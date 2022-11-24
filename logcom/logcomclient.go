package logcom

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom/log"
	"github.com/google/uuid"
)

var (
	configuration     Configuration
	apiClientInstance *logcomapi.APIClient
	internalLogger    zerolog.Logger
	once              sync.Once
)

type ClientCredentialProvider interface {
	GetClientCredential() (string, error)
}

type Configuration struct {
	ServiceName              string
	LogComURL                string
	HeaderProvider           HeaderProviderFunc
	ClientCredentialProvider ClientCredentialProvider
}

type HeaderProviderFunc func(ctx context.Context) http.Header

func init() {
	internalLogger = log.Logger

	logcomURL := os.Getenv("LOG_COM_URL")
	if logcomURL == "" {
		internalLogger.Error().Msg("LogCom URL is missing thus functionalities are not available")
		return
	}

	config := Configuration{
		ServiceName: "Unknown",
		LogComURL:   logcomURL,
		HeaderProvider: func(ctx context.Context) http.Header {
			headers := make(map[string][]string, 0)

			if authorization := ctx.Value("Authorization").(*string); authorization != nil || *authorization != "" {
				headers["Authorization"] = []string{*authorization}
			} else {
				internalLogger.Warn().Msg("Authorization header not set")
			}

			if requestID := ctx.Value("X-Request-ID").(*string); requestID != nil || *requestID != "" {
				headers["X-Request-ID"] = []string{*requestID}
			} else {
				internalLogger.Warn().Msg("X-Request-ID header not set")
			}

			return headers
		},
	}

	configuration = config

	logComAPIConfig := logcomapi.NewConfiguration()
	logComAPIConfig.BasePath = configuration.LogComURL + "/api"
	logComAPIConfig.Host = configuration.LogComURL

	parsedUrl, err := url.Parse(configuration.LogComURL)
	if err != nil {
		internalLogger.Error().Msgf("Failed to get LogCom URL scheme, falling back to default (%s)", "https")

		parsedUrl = &url.URL{
			Scheme: "https",
		}
	}

	logComAPIConfig.Scheme = parsedUrl.Scheme

	apiClientInstance = logcomapi.NewAPIClient(logComAPIConfig)
}

func Init(config Configuration, logger *zerolog.Logger) {
	if config.LogComURL == "" {
		logcomURL := os.Getenv("LOG_COM_URL")
		if logcomURL == "" {
			log.Fatal().Msg("LogCom URL is missing")
			return
		}
	}

	once.Do(func() {
		configuration = config

		logcomAPIConfig := logcomapi.NewConfiguration()
		logcomAPIConfig.BasePath = configuration.LogComURL + "/api"
		logcomAPIConfig.Host = configuration.LogComURL

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

		logcomAPIConfig.Scheme = parsedUrl.Scheme

		apiClientInstance = logcomapi.NewAPIClient(logcomAPIConfig)
		*logger = logger.Hook(logComHook{internalLogger: internalLogger})
	})
}

func SetClientCredentialProvider(provider ClientCredentialProvider) {
	configuration.ClientCredentialProvider = provider
}

func IsEnabled() bool {
	return apiClientInstance != nil
}

func SendConsoleLog(ctx context.Context, logLevel zerolog.Level, message string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendConsoleLog(ctx, logLevel, message, nil)
}

func SendConsoleLogWithModel(ctx context.Context, model logcomapi.CreateConsoleLogRequestDto) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendConsoleLogWithModel(ctx, model, nil)
}

func SendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{}) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithCreation(ctx, subject, subjectName, newValue, nil)
}

func SendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{}, ignoredProperties ...string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithModification(ctx, subject, subjectName, oldValue, newValue, nil, ignoredProperties...)
}

func SendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithModificationModelChanges(ctx, subject, subjectName, changes, nil)
}

func SendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{}) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLogWithDeletion(ctx, subject, subjectName, oldValue, nil)
}

func SendAuditLog(ctx context.Context, model logcomapi.CreateAuditLogRequestDto) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLog(ctx, model, nil)
}

func SendAuditLogGroup(ctx context.Context, auditLogCollector *AuditLogCollector) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendAuditLog(ctx, auditLogCollector.get(), nil)
}

func SendNotification(ctx context.Context, eventCategory string, message string, targets map[string][]string) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendNotification(ctx, eventCategory, message, targets, nil)
}

func SendNotificationWithModel(ctx context.Context, model logcomapi.CreateNotificationRequestDto) error {
	if ctx == context.TODO() || ctx == context.Background() {
		return errors.New("context cannot be empty")
	}
	return sendNotificationWithModel(ctx, model, nil)
}

func sendConsoleLog(ctx context.Context, logLevel zerolog.Level, message string, extraHeaders http.Header) error {
	return sendConsoleLogWithModel(ctx, logcomapi.CreateConsoleLogRequestDto{
		Level:   int32(logLevel),
		Message: message,
		Service: configuration.ServiceName,
	}, extraHeaders)
}

func sendConsoleLogWithModel(ctx context.Context, model logcomapi.CreateConsoleLogRequestDto, extraHeaders http.Header) error {
	if !IsEnabled() {
		internalLogger.Debug().Msg("LogCom is disabled")
		return nil
	}

	prepareConsoleLogRequestDTO(&model)

	headers := configuration.HeaderProvider(ctx)
	for headerName, headerValue := range extraHeaders {
		headers[headerName] = headerValue
	}

	result, err := apiClientInstance.ConsoleLogApi.CreateConsoleLogV1Int(ctx, model, requestConfigurer(ctx, headers))
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func sendAuditLogWithCreation(ctx context.Context, subject, subjectName string, newValue interface{}, extraHeaders http.Header) error {
	if newValue == nil {
		newValue = ""
	}
	return sendAuditLog(ctx, logcomapi.CreateAuditLogRequestDto{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     subject,
		SubjectName: subjectName,
	}, extraHeaders)
}

func sendAuditLogWithModification(ctx context.Context, subject, subjectName string, oldValue, newValue interface{}, extraHeaders http.Header, ignoredProperties ...string) error {
	if isPrimitiveType(oldValue) {
		return SendAuditLog(ctx, logcomapi.CreateAuditLogRequestDto{
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

	return sendAuditLogWithModificationModelChanges(ctx, subject, subjectName, changes, extraHeaders)
}

func sendAuditLogWithModificationModelChanges(ctx context.Context, subject, subjectName string, changes []ModelChange, extraHeaders http.Header) error {
	dto := logcomapi.CreateAuditLogRequestDto{
		Category:    "MODIFICATION",
		Subject:     subject,
		SubjectName: subjectName,
	}

	if changesCount := len(changes); changesCount > 1 {
		changesDTO := make([]logcomapi.NewAuditLogChangeDto, changesCount)

		for i, change := range changes {
			changesDTO[i] = logcomapi.NewAuditLogChangeDto{
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

	return sendAuditLog(ctx, dto, extraHeaders)
}

func sendAuditLogWithDeletion(ctx context.Context, subject, subjectName string, oldValue interface{}, extraHeaders http.Header) error {
	if oldValue == nil {
		oldValue = ""
	}
	return sendAuditLog(ctx, logcomapi.CreateAuditLogRequestDto{
		Category:    "DELETION",
		OldValue:    stringify(oldValue),
		Subject:     subject,
		SubjectName: subjectName,
	}, extraHeaders)
}

func sendAuditLog(ctx context.Context, model logcomapi.CreateAuditLogRequestDto, extraHeaders http.Header) error {
	if !IsEnabled() {
		internalLogger.Info().Msg("LogCom is disabled")
		return nil
	}

	prepareAuditLogRequestDTO(&model)

	headers := configuration.HeaderProvider(ctx)
	for headerName, headerValue := range extraHeaders {
		headers[headerName] = headerValue
	}

	result, err := apiClientInstance.AuditLogApi.CreateAuditLogV1Int(ctx, model, requestConfigurer(ctx, headers))
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func sendAuditLogGroup[T any](ctx context.Context, auditLog *auditLog[T], auditLogCollector *AuditLogCollector, extraHeaders http.Header) error {
	auditLogCollector.parentAuditLog = logcomapi.CreateAuditLogRequestDto{
		Category:    auditLog.operation,
		Subject:     auditLog.subject,
		SubjectName: auditLog.subjectName,
	}
	return sendAuditLog(ctx, auditLogCollector.get(), extraHeaders)
}

func sendNotification(ctx context.Context, eventCategory string, message string, targets map[string][]string, extraHeaders http.Header) error {
	return sendNotificationWithModel(ctx, logcomapi.CreateNotificationRequestDto{
		EventCategory: eventCategory,
		Message:       message,
		Service:       configuration.ServiceName,
		Targets:       targets,
	}, extraHeaders)
}

func sendNotificationWithModel(ctx context.Context, model logcomapi.CreateNotificationRequestDto, extraHeaders http.Header) error {
	if !IsEnabled() {
		internalLogger.Debug().Msg("LogCom is disabled")
		return nil
	}

	headers := configuration.HeaderProvider(ctx)
	for headerName, headerValue := range extraHeaders {
		headers[headerName] = headerValue
	}

	prepareNotificationRequestDTO(&model)

	result, err := apiClientInstance.NotificationApi.CreateNotificationV1(ctx, model, requestConfigurer(ctx, headers))
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
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
			} else {
				requestID = uuid.NewString()
			}
		}

		request.Header.Add("X-Request-ID", requestID)

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

func ensureConsoleLog(consoleLogPointer **consoleLog) {
	if *consoleLogPointer == nil {
		*consoleLogPointer = &consoleLog{}
	}
}

func ensureNotification[T any](notificationPointer **notification[T]) {
	if *notificationPointer == nil {
		*notificationPointer = &notification[T]{}
	}
}

func ensureHTTPHeaders(httpHeadersPointer *http.Header) {
	if *httpHeadersPointer == nil {
		*httpHeadersPointer = make(map[string][]string, 0)
	}
}
