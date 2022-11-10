package logcom

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/logcom-api"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom/log"
	"github.com/google/uuid"
)

type AuditLogAction interface {
	IgnoreChangeOf(propertyName string) AuditLogAction
	AndNotify() NotificationConfigurer[AuditLogAction]
	AndNotifyRoles(targets ...string) AuditLogAction
	AndNotifySessions(targets ...string) AuditLogAction
	AndNotifyUsers(targets ...string) AuditLogAction
	AndLog(logLevel zerolog.Level, message string) AuditLogAction
	Send() error
}

type AuditLogInitializer interface {
	WithClientSecret(secret string) AuditLogAction
	WithBearerAuthorization(bearerToken string) AuditLogAction
	WithContext(ctx context.Context) AuditLogAction
	WithTransactionID(transactionID uuid.UUID) AuditLogInitializer
}

type AuditLogOperation interface {
	Create(subject, subjectName string, newValue interface{}) AuditLogInitializer
	Delete(subject, subjectName string, oldValue interface{}) AuditLogInitializer
	Modify(subject, subjectName string, oldValue interface{}, newValue interface{}) AuditLogInitializer
}

type AuditLogModelDiff interface {
	GetChanges(model interface{}, ignoredProperties map[string]interface{}) []ModelChange
}

type ModelChange struct {
	PropertyName string
	OldValue     interface{}
	NewValue     interface{}
}

type AuditLogCollector struct {
	parentAuditLog logcomapi.CreateAuditLogRequestDto
	auditLogs      map[string]map[string][]logcomapi.CreateAuditLogRequestDto
}

type auditLog[T any] struct {
	ctx               context.Context
	operation         string
	newValue          interface{}
	oldValue          interface{}
	subject           string
	subjectName       string
	ignoredProperties []string
	httpHeaders       http.Header
	consoleLog        *consoleLog
	notification      *notification[T]
}

func Audit() AuditLogOperation {
	return &auditLog[AuditLogAction]{}
}

func AuditCreation(subject, subjectName string, newValue interface{}) AuditLogInitializer {
	return &auditLog[AuditLogAction]{
		operation:   "CREATION",
		newValue:    newValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func AuditModification(subject, subjectName string, oldValue, newValue interface{}) AuditLogInitializer {
	return &auditLog[AuditLogAction]{
		operation:   "MODIFICATION",
		oldValue:    oldValue,
		newValue:    newValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func AuditDeletion(subject, subjectName string, oldValue interface{}) AuditLogInitializer {
	return &auditLog[AuditLogAction]{
		operation:   "DELETION",
		oldValue:    oldValue,
		subject:     subject,
		subjectName: subjectName,
	}
}

func NewAuditLogCollector(parentSubject, parentSubjectName string) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: logcomapi.CreateAuditLogRequestDto{
			Subject:     parentSubject,
			SubjectName: parentSubjectName},
		auditLogs: make(map[string]map[string][]logcomapi.CreateAuditLogRequestDto, 0),
	}
}

func NewAuditLogCollectorWithParent(parentAuditLog logcomapi.CreateAuditLogRequestDto) *AuditLogCollector {
	return &AuditLogCollector{
		parentAuditLog: parentAuditLog,
		auditLogs:      make(map[string]map[string][]logcomapi.CreateAuditLogRequestDto, 0),
	}
}

func GetModelChanges(ctx context.Context, oldModel, newModel interface{}, ignoredProperties ...string) ([]ModelChange, error) {
	ignoredPropertySet := make(map[string]struct{}, len(ignoredProperties))
	for i := range ignoredProperties {
		ignoredPropertySet[ignoredProperties[i]] = struct{}{}
	}

	valueOfOldModel := reflect.ValueOf(oldModel)
	typeOfOldModel := valueOfOldModel.Type()
	kindOfOldModel := typeOfOldModel.Kind()
	valueOfNewModel := reflect.ValueOf(newModel)
	typeOfNewModel := valueOfNewModel.Type()
	var fieldCountOfOldModel int

	if kindOfOldModel == reflect.Ptr {
		fieldCountOfOldModel = valueOfOldModel.Elem().NumField()
	} else {
		fieldCountOfOldModel = valueOfOldModel.NumField()
	}

	if kindOfOldModel == reflect.Func || kindOfOldModel == reflect.Chan {
		return nil, errors.New("unsupported type: " + kindOfOldModel.String())
	}

	if typeOfOldModel != typeOfNewModel {
		log.Fatal().MsgContext(ctx, "Old and new model have different types")
		return nil, errors.New("old and new model have different types")
	}

	changes := make([]ModelChange, 0)
	for i := 0; i < fieldCountOfOldModel; i++ {
		if valueOfOldModel.Field(i).CanInterface() {
			fieldNameOfOldModel := typeOfOldModel.Field(i).Name

			if _, ok := ignoredPropertySet[fieldNameOfOldModel]; ok {
				continue
			}

			fieldOfOldModel := valueOfOldModel.Field(i)
			valueOfOldModelField := fieldOfOldModel.Interface()
			valueOfNewModelField := valueOfNewModel.FieldByName(fieldNameOfOldModel).Interface()

			kindOfOldModelField := fieldOfOldModel.Kind()
			if (kindOfOldModelField == reflect.Func) || (kindOfOldModelField == reflect.Chan) {
				log.Warn().MsgfContext(ctx, "Unsupported type: %s", kindOfOldModelField.String())
				continue
			}

			if (kindOfOldModelField >= reflect.Array) && (kindOfOldModelField != reflect.String) {
				if !reflect.DeepEqual(valueOfOldModelField, valueOfNewModelField) {
					changes = append(changes, ModelChange{
						PropertyName: fieldNameOfOldModel,
						OldValue:     valueOfOldModelField,
						NewValue:     valueOfNewModelField,
					})
				}
				continue
			}

			if valueOfOldModelField != valueOfNewModelField {
				changes = append(changes, ModelChange{
					PropertyName: fieldNameOfOldModel,
					OldValue:     valueOfOldModelField,
					NewValue:     valueOfNewModelField,
				})
			}
		}
	}

	return changes, nil
}

func prepareAuditLogRequestDTO(dto *logcomapi.CreateAuditLogRequestDto) {
	if dto.ServiceCreated == "" {
		dto.ServiceCreated = configuration.ServiceName
	}

	if dto.ServiceAffected == "" {
		dto.ServiceAffected = configuration.ServiceName
	}

	if dto.CreatedAt != nil {
		utcNow := dto.CreatedAt.UTC()
		dto.CreatedAt = &utcNow
	} else {
		utcNow := time.Now().UTC()
		dto.CreatedAt = &utcNow
	}
}

func (al *auditLog[T]) IgnoreChangeOf(propertyName string) AuditLogAction {
	al.ensureIgnoredProperties()
	al.ignoredProperties = append(al.ignoredProperties, propertyName)
	return al
}

func (al *auditLog[T]) AndNotify() NotificationConfigurer[AuditLogAction] {
	ensureNotification(&al.notification)
	al.notification.eventCategory = "NOTIFICATION"
	return interface{}(al).(NotificationConfigurer[AuditLogAction])
}

func (al *auditLog[T]) AndNotifyRoles(targets ...string) AuditLogAction {
	ensureNotification(&al.notification)
	al.notification.Roles(targets...)
	return al
}

func (al *auditLog[T]) AndNotifySessions(targets ...string) AuditLogAction {
	ensureNotification(&al.notification)
	al.notification.Sessions(targets...)
	return al
}

func (al *auditLog[T]) AndNotifyUsers(targets ...string) AuditLogAction {
	ensureNotification(&al.notification)
	al.notification.Users(targets...)
	return al
}

func (al *auditLog[T]) AndLog(logLevel zerolog.Level, message string) AuditLogAction {
	ensureConsoleLog(&al.consoleLog)
	al.consoleLog.logLevel = logLevel
	al.consoleLog.message = message
	return al
}

func (al *auditLog[T]) WithClientSecret(secret string) AuditLogAction {
	ensureHTTPHeaders(&al.httpHeaders)
	al.httpHeaders["X-Client-Secret"] = []string{secret}
	return al
}

func (al *auditLog[T]) WithBearerAuthorization(bearerToken string) AuditLogAction {
	ensureHTTPHeaders(&al.httpHeaders)
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		bearerToken = "Bearer " + bearerToken
	}
	al.httpHeaders["Authorization"] = []string{bearerToken}
	return al
}

func (al *auditLog[T]) WithContext(ctx context.Context) AuditLogAction {
	al.ctx = ctx
	return al
}

func (al *auditLog[T]) WithTransactionID(transactionID uuid.UUID) AuditLogInitializer {
	ensureHTTPHeaders(&al.httpHeaders)
	al.httpHeaders["X-Request-ID"] = []string{transactionID.String()}
	return al
}

func (al *auditLog[T]) Send() error {
	var err error

	switch al.operation {
	case "CREATION":
		err = SendAuditLogWithCreation(al.ctx, al.subject, al.subjectName, al.newValue)
	case "MODIFICATION":
		err = SendAuditLogWithModification(al.ctx, al.subject, al.subjectName, al.oldValue, al.newValue, al.ignoredProperties...)
	case "DELETION":
		err = SendAuditLogWithDeletion(al.ctx, al.subject, al.subjectName, al.oldValue)
	default:
		return errors.New("invalid audit operation: " + al.operation)
	}

	if err != nil {
		log.Error().Msg("Failed to send audit log")
		return err
	}

	if al.consoleLog != nil {
		if err = SendNotification(al.ctx, al.notification.eventCategory, al.notification.message, al.notification.targets); err != nil {
			log.Error().Err(err).Msg("Failed to send notification")
		}
	}
	if al.notification != nil {
		if err = SendConsoleLog(al.ctx, al.consoleLog.logLevel, al.consoleLog.message); err != nil {
			log.Error().Err(err).Msg("Failed to send console log")
			// Todo: Maybe add it to a list for retry purpose
		}
	}
	return nil
}

func (al *auditLog[T]) Create(subject, subjectName string, newValue interface{}) AuditLogInitializer {
	al.operation = "CREATION"
	al.newValue = newValue
	al.subject = subject
	al.subjectName = subjectName
	return al
}

func (al *auditLog[T]) Modify(subject, subjectName string, oldValue, newValue interface{}) AuditLogInitializer {
	al.operation = "MODIFICATION"
	al.oldValue = oldValue
	al.newValue = newValue
	al.subject = subject
	al.subjectName = subjectName
	return al
}

func (al *auditLog[T]) Delete(subject, subjectName string, oldValue interface{}) AuditLogInitializer {
	al.operation = "DELETION"
	al.oldValue = oldValue
	al.subject = subject
	al.subjectName = subjectName
	return al
}

func (al *auditLog[T]) Roles(targets ...string) NotificationConfigurer[T] {
	al.notification.Roles(targets...)
	return al
}

func (al *auditLog[T]) Sessions(targets ...string) NotificationConfigurer[T] {
	al.notification.Sessions(targets...)
	return al
}

func (al *auditLog[T]) Users(targets ...string) NotificationConfigurer[T] {
	al.notification.Users(targets...)
	return al
}

func (al *auditLog[T]) Message(message string) T {
	al.notification.Message(message)
	return interface{}(al).(T)
}

func (al *auditLog[T]) ensureIgnoredProperties() {
	if al.ignoredProperties == nil {
		al.ignoredProperties = make([]string, 0)
	}
}

func (c *AuditLogCollector) AddCreation(itemSubject, itemSubjectName string, newValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDto{
		Category:    "CREATION",
		NewValue:    stringify(newValue),
		Subject:     itemSubject,
		SubjectName: itemSubjectName,
	})
}

func (c *AuditLogCollector) AddModification(itemSubject, itemSubjectName string, oldValue, newValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDto{
		Category:    "MODIFICATION",
		NewValue:    stringify(newValue),
		OldValue:    stringify(oldValue),
		Subject:     itemSubject,
		SubjectName: itemSubjectName,
	})
}

func (c *AuditLogCollector) AddDeletion(itemSubject, itemSubjectName string, oldValue interface{}) {
	c.Add(logcomapi.CreateAuditLogRequestDto{
		Category:    "DELETION",
		OldValue:    stringify(oldValue),
		Subject:     itemSubject,
		SubjectName: itemSubjectName,
	})
}

func (c *AuditLogCollector) Add(model logcomapi.CreateAuditLogRequestDto) {
	if subjectNameMap, ok := c.auditLogs[model.Subject]; ok {
		if auditLogList, ok := subjectNameMap[model.SubjectName]; ok {
			c.auditLogs[model.Subject][model.SubjectName] = append(auditLogList, model)
		} else {
			c.auditLogs[model.Subject][model.SubjectName] = []logcomapi.CreateAuditLogRequestDto{model}
		}
	} else {
		c.auditLogs[model.Subject] = map[string][]logcomapi.CreateAuditLogRequestDto{model.SubjectName: {model}}
	}
}

func (c *AuditLogCollector) get() logcomapi.CreateAuditLogRequestDto {
	c.parentAuditLog.GroupedChanges = make([]logcomapi.NewAuditLogChangeDto, 0)
	hasSameCategory := true
	seenCategory := ""

	for subjectAsKey, subjectNameGroupAsValue := range c.auditLogs {
		for subjectNameAsKey, auditLogList := range subjectNameGroupAsValue {
			for _, auditLog := range auditLogList {
				c.parentAuditLog.GroupedChanges = append(c.parentAuditLog.GroupedChanges, logcomapi.NewAuditLogChangeDto{
					Category:            auditLog.Category,
					Message:             auditLog.Message,
					NewValue:            auditLog.NewValue,
					OldValue:            auditLog.OldValue,
					Subject:             subjectAsKey,
					SubjectName:         subjectNameAsKey,
					SubjectPropertyName: auditLog.SubjectPropertyName,
				})

				if hasSameCategory && seenCategory != auditLog.Category && seenCategory != "" {
					hasSameCategory = false
				}

				seenCategory = auditLog.Category
			}
		}
	}

	if hasSameCategory {
		c.parentAuditLog.Category = seenCategory
	} else {
		c.parentAuditLog.Category = "MODIFICATION"
	}

	return c.parentAuditLog
}
