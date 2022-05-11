package logcomapi

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
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

type logComHook struct {
	internalLogger zerolog.Logger
}

func (h logComHook) Run(e *zerolog.Event, ctx context.Context, level zerolog.Level, msg string) {
	if err := SendConsoleLog(ctx, level, msg); err != nil {
		h.internalLogger.Error().Err(err).Msg("Failed to send console log to LogCom")
	}
}

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

	if model.Service == "" {
		model.Service = configuration.ServiceName
	}

	if model.CreatedAt != nil {
		now := time.Now().UTC()
		model.CreatedAt = &now
	} else {
		now := time.Now().UTC()
		model.CreatedAt = &now
	}

	headers := configuration.HeaderProvider(ctx)

	result, err := instance.ConsoleLogApi.CreateConsoleLogV1Int(ctx, model, func(request *http.Request) {
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
	})
	if err != nil {
		return err
	}

	if !isHTTPStatusSuccess(result.StatusCode) {
		return errors.New(result.Status)
	}

	return nil
}

func SendAuditLog(ctx context.Context, model CreateAuditLogRequestDto) error {
	if !IsEnabled() {
		internalLogger.Debug().Msg("LogCom is disabled")
		return nil
	}

	result, err := instance.AuditLogApi.CreateAuditLogV1Int(ctx, model, NilRequestConfigurer)
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
