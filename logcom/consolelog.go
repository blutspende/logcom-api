package logcom

import (
	"context"
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

type consoleLog struct {
	ctx      context.Context
	logLevel Level
	message  string
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
