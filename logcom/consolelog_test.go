package logcom

import (
	"context"
	"testing"

	logcomapi "github.com/blutspende/logcom-api"
)

func TestSend(t *testing.T) {
	ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")

	err := Log(ctx).
		Level(logcomapi.Debug).
		Message("Debug message").
		Build().
		Send()

	if err != nil {
		t.Errorf("expected no error")
	}
}
