package logcom

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom/log"
	"github.com/google/uuid"
)

type NotificationAction interface {
	AndLog(logLevel zerolog.Level, message string) NotificationAction
	Send() error
}

type NotificationInitializer interface {
	UseService2ServiceAuthorization() NotificationAction
	WithBearerAuthorization(bearerToken string) NotificationAction
	WithContext(ctx context.Context) NotificationAction
	WithTransactionID(transactionID uuid.UUID) NotificationInitializer
}

type NotificationConfigurer[T any] interface {
	Message(message string) T
	Roles(targets ...string) NotificationConfigurer[T]
	Sessions(targets ...string) NotificationConfigurer[T]
	Users(targets ...string) NotificationConfigurer[T]
}

type NotificationMessage[T any] interface {
	Message(message string) T
}

type notification[T any] struct {
	ctx           context.Context
	httpHeaders   http.Header
	eventCategory string
	message       string
	targets       map[string][]string
	consoleLog    *consoleLog
}

func Notify() NotificationConfigurer[NotificationInitializer] {
	return &notification[NotificationInitializer]{
		eventCategory: "NOTIFICATION",
		ctx:           context.TODO(),
	}
}

func prepareNotificationRequestDTO(dto *logcomapi.CreateNotificationRequestDto) {
	if dto.Service == "" {
		dto.Service = configuration.ServiceName
	}

	if dto.EventCategory == "" {
		dto.EventCategory = "NOTIFICATION"
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}

func (n *notification[T]) AndLog(logLevel zerolog.Level, message string) NotificationAction {
	ensureConsoleLog(&n.consoleLog)
	n.consoleLog.logLevel = logLevel
	n.consoleLog.message = message
	return n
}

func (n *notification[T]) Send() error {
	if err := sendNotification(n.ctx, n.eventCategory, n.message, n.targets, n.httpHeaders); err != nil {
		log.Error().Msg("Failed to send notification")
		return err
	}

	if n.consoleLog != nil {
		if err := sendConsoleLog(n.ctx, n.consoleLog.logLevel, n.consoleLog.message, n.httpHeaders); err != nil {
			log.Error().Err(err).Msg("Failed to send console log")
		}
	}

	return nil
}

func (n *notification[T]) UseService2ServiceAuthorization() NotificationAction {
	if configuration.ClientCredentialProvider == nil {
		log.Fatal().Msg("Client credential provider must be set")
		return n
	}

	clientCredential, err := configuration.ClientCredentialProvider.GetClientCredential()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get client credential")
		return n
	}

	return n.WithBearerAuthorization(clientCredential)
}

func (n *notification[T]) WithBearerAuthorization(bearerToken string) NotificationAction {
	ensureHTTPHeaders(&n.httpHeaders)
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	n.httpHeaders["Authorization"] = []string{bearerToken}
	return n
}

func (n *notification[T]) WithContext(ctx context.Context) NotificationAction {
	n.ctx = ctx
	return n
}

func (n *notification[T]) WithTransactionID(transactionID uuid.UUID) NotificationInitializer {
	ensureHTTPHeaders(&n.httpHeaders)
	n.httpHeaders["X-Request-ID"] = []string{transactionID.String()}
	return n
}

func (n *notification[T]) Roles(targets ...string) NotificationConfigurer[T] {
	n.transformToNotificationTargets("ROLE", targets...)
	return n
}

func (n *notification[T]) Sessions(targets ...string) NotificationConfigurer[T] {
	n.transformToNotificationTargets("SESSION", targets...)
	return n
}

func (n *notification[T]) Users(targets ...string) NotificationConfigurer[T] {
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
