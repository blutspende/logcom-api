package logcom

import (
	"context"
	"testing"

	logcomapi "github.com/blutspende/logcom-api"
)

func TestSendAndLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")

	err := Notify(ctx).
		Roles("test_role").
		Message("Test notification").
		Build().
		AndLog(logcomapi.Debug, "Debug log").
		Send()

	if err != nil {
		t.Errorf("expected no error")
	}
}
