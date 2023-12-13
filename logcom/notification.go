package logcom

import (
	"context"
	"strings"
	"time"

	"github.com/blutspende/logcom-api"
	"github.com/google/uuid"
)

type NotificationAction interface {
	AndLog(logLevel logcomapi.LogLevel, message string) NotificationAction
	OnComplete(onCompleteCallback func(error)) NotificationAction
	Send() error
}

type NotificationConfiguration interface {
	UseService2ServiceAuthorization() NotificationAction
	WithBearerAuthorization(bearerToken string) NotificationAction
	WithContext(ctx context.Context) NotificationAction
	WithTransactionID(transactionID uuid.UUID) NotificationConfiguration
}

type NotificationOperation[T any] interface {
	Roles(targets ...string) NotificationMessage[T]
	Sessions(targets ...string) NotificationMessage[T]
	Users(targets ...string) NotificationMessage[T]
}

type NotificationMessage[T any] interface {
	NotificationOperation[T]
	Message(message string) T
}

type notification[T any] struct {
	ctx                context.Context
	eventCategory      logcomapi.NotificationEventCategory
	message            string
	targets            map[string][]string
	consoleLog         *consoleLog
	onCompleteCallback func(error)
}

func Notify() NotificationOperation[NotificationConfiguration] {
	return &notification[NotificationConfiguration]{
		eventCategory:      logcomapi.Notification,
		ctx:                context.TODO(),
		onCompleteCallback: func(error) {},
	}
}

func prepareNotificationRequestDTO(dto *logcomapi.CreateNotificationRequestDTO) {
	if dto.Service == "" {
		dto.Service = configuration.ServiceName
	}

	if dto.EventCategory == nil {
		dto.EventCategory = toPtr(logcomapi.Notification)
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}

func (n *notification[T]) AndLog(logLevel logcomapi.LogLevel, message string) NotificationAction {
	ensureConsoleLog(&n.consoleLog)
	n.consoleLog.logLevel = logLevel
	n.consoleLog.message = message
	return n
}

func (n *notification[T]) OnComplete(onCompleteCallback func(error)) NotificationAction {
	n.onCompleteCallback = onCompleteCallback
	return n
}

func (n *notification[T]) Send() error {
	err := sendNotification(n.ctx, n.eventCategory, n.message, n.targets)
	if err != nil {
		logError.Println("Failed to send notification")
		n.onCompleteCallback(err)
		return err
	}

	if n.consoleLog != nil {
		if err = sendConsoleLog(n.ctx, n.consoleLog.logLevel, n.consoleLog.message); err != nil {
			logError.Printf("Failed to send console log: %v\n", err)
		}
	}

	n.onCompleteCallback(err)

	return nil
}

func (n *notification[T]) UseService2ServiceAuthorization() NotificationAction {
	if configuration.ClientCredentialProvider == nil {
		logFatal.Println("Client credential provider must be set")
		return n
	}

	clientCredential, err := configuration.ClientCredentialProvider.GetClientCredential()
	if err != nil {
		logError.Printf("Failed to get client credential: %v\n", err)
		return n
	}

	return n.WithBearerAuthorization(clientCredential)
}

func (n *notification[T]) WithBearerAuthorization(bearerToken string) NotificationAction {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	n.ctx = context.WithValue(n.ctx, "Authorization", bearerToken)
	return n
}

func (n *notification[T]) WithContext(ctx context.Context) NotificationAction {
	n.ctx = ctx
	return n
}

func (n *notification[T]) WithTransactionID(transactionID uuid.UUID) NotificationConfiguration {
	n.ctx = context.WithValue(n.ctx, "RequestID", transactionID.String())
	return n
}

func (n *notification[T]) Roles(targets ...string) NotificationMessage[T] {
	n.transformToNotificationTargets("ROLE", targets...)
	return n
}

func (n *notification[T]) Sessions(targets ...string) NotificationMessage[T] {
	n.transformToNotificationTargets("SESSION", targets...)
	return n
}

func (n *notification[T]) Users(targets ...string) NotificationMessage[T] {
	n.transformToNotificationTargets("USER", targets...)
	return n
}

func (n *notification[T]) Message(message string) T {
	n.message = message
	return interface{}(n).(T)
}

func (n *notification[T]) transformToNotificationTargets(targetType string, targets ...string) {
	if n.targets == nil {
		n.targets = make(map[string][]string, 0)
	}

	n.targets[targetType] = make([]string, len(targets))
	for i, target := range targets {
		n.targets[targetType][i] = target
	}
}
