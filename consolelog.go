package logcomapi

import (
	"context"
	"time"

	"github.com/DRK-Blutspende-BaWueHe/zerolog-for-logcom"
)

type logComHook struct {
	internalLogger zerolog.Logger
}

func (h logComHook) Run(e *zerolog.Event, ctx context.Context, level zerolog.Level, msg string) {
	if err := SendConsoleLog(ctx, level, msg); err != nil {
		h.internalLogger.Error().Err(err).Msg("Failed to send console log to LogCom")
	}
}

func prepareConsoleLogRequestDTO(dto *CreateConsoleLogRequestDto) {
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
