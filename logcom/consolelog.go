package logcom

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/blutspende/logcom-api"
	"github.com/google/uuid"
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
	Level(level logcomapi.LogLevel) ConsoleLogOperation
	Message(message string) ConsoleLogConfiguration
	MessageF(format string, params ...any) ConsoleLogConfiguration
}

type consoleLog struct {
	ctx                context.Context
	logLevel           logcomapi.LogLevel
	message            string
	onCompleteCallback func(error)
}

func Log() ConsoleLogOperation {
	return &consoleLog{
		ctx:                context.TODO(),
		logLevel:           logcomapi.Warning,
		onCompleteCallback: func(error) {},
	}
}

func prepareConsoleLogRequestDTO(dto *logcomapi.CreateConsoleLogRequestDTO) {
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

func (cl *consoleLog) Level(level logcomapi.LogLevel) ConsoleLogOperation {
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
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	cl.ctx = context.WithValue(cl.ctx, "Authorization", bearerToken)
	return cl
}

func (cl *consoleLog) WithContext(ctx context.Context) ConsoleLogAction {
	cl.ctx = ctx
	return cl
}

func (cl *consoleLog) WithTransactionID(transactionID uuid.UUID) ConsoleLogConfiguration {
	cl.ctx = context.WithValue(cl.ctx, "RequestID", transactionID.String())
	return cl
}

func (cl *consoleLog) OnComplete(onCompleteCallback func(error)) ConsoleLogAction {
	cl.onCompleteCallback = onCompleteCallback
	return cl
}

func (cl *consoleLog) Send() error {
	err := sendConsoleLog(cl.ctx, cl.logLevel, cl.message)
	if err != nil {
		logError.Printf("Failed to send console log: %v\n", err)
	}

	cl.onCompleteCallback(err)

	return err
}
