package logcom

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
)

type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
)

type ConsoleLogAction interface {
	OnComplete(onCompleteCallback func(error)) ConsoleLogAction
	Send() error
}

type ConsoleLogConfiguration interface {
	UseService2ServiceAuthorization() ConsoleLogAction
	WithBearerAuthorization(bearerToken string) ConsoleLogAction
	WithContext(ctx context.Context) ConsoleLogAction
	WithTransactionID(transactionID uuid.UUID) ConsoleLogConfiguration
}

type ConsoleLogOperation interface {
	Level(level Level) ConsoleLogOperation
	Message(message string) ConsoleLogConfiguration
	MessageF(format string, params ...any) ConsoleLogConfiguration
}

type consoleLog struct {
	ctx                context.Context
	logLevel           Level
	message            string
	httpHeaders        http.Header
	onCompleteCallback func(error)
}

func Log() ConsoleLogOperation {
	return &consoleLog{
		ctx:                context.TODO(),
		logLevel:           WarnLevel,
		onCompleteCallback: func(error) {},
	}
}

func prepareConsoleLogRequestDTO(dto *logcomapi.CreateConsoleLogRequestDto) {
	if dto.Service == "" {
		dto.Service = configuration.ServiceName
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}

func (cl *consoleLog) Level(level Level) ConsoleLogOperation {
	cl.logLevel = level
	return cl
}

func (cl *consoleLog) Message(message string) ConsoleLogConfiguration {
	cl.message = message
	return cl
}

func (cl *consoleLog) MessageF(format string, params ...any) ConsoleLogConfiguration {
	cl.message = fmt.Sprintf(format, params...)
	return cl
}

func (cl *consoleLog) UseService2ServiceAuthorization() ConsoleLogAction {
	if configuration.ClientCredentialProvider == nil {
		logFatal.Println("Client credential provider must be set")
		return cl
	}

	clientCredential, err := configuration.ClientCredentialProvider.GetClientCredential()
	if err != nil {
		logError.Printf("Failed to get client credential: %v\n", err)
		return cl
	}

	return cl.WithBearerAuthorization(clientCredential)
}

func (cl *consoleLog) WithBearerAuthorization(bearerToken string) ConsoleLogAction {
	ensureHTTPHeaders(&cl.httpHeaders)
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	cl.httpHeaders["Authorization"] = []string{bearerToken}
	return cl
}

func (cl *consoleLog) WithContext(ctx context.Context) ConsoleLogAction {
	cl.ctx = ctx
	return cl
}

func (cl *consoleLog) WithTransactionID(transactionID uuid.UUID) ConsoleLogConfiguration {
	ensureHTTPHeaders(&cl.httpHeaders)
	cl.httpHeaders["X-Request-ID"] = []string{transactionID.String()}
	return cl
}

func (cl *consoleLog) OnComplete(onCompleteCallback func(error)) ConsoleLogAction {
	cl.onCompleteCallback = onCompleteCallback
	return cl
}

func (cl *consoleLog) Send() error {
	err := sendConsoleLog(cl.ctx, cl.logLevel, cl.message, cl.httpHeaders)
	if err != nil {
		logError.Printf("Failed to send console log: %v\n", err)
	}

	cl.onCompleteCallback(err)

	return err
}
