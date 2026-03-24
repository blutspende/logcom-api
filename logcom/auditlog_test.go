package logcom

import (
	"context"
	"testing"

	logcomapi "github.com/blutspende/logcom-api"
	"github.com/google/uuid"
)

func TestSendAndNotifyAndLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")
	ctx = context.WithValue(ctx, "RequestID", uuid.NewString())

	err := Audit(ctx).
		Create("SUBJECT", "NAME", nil).
		WithTransactionID(uuid.New()).
		WithBearerAuthorization("BearerToken").
		Build().
		AndNotify().
		Roles("test_role").
		Message("Test notification").
		AndLog(logcomapi.Debug, "Debug log").
		Send()

	if err != nil {
		t.Errorf("expected no error")
	}
}
