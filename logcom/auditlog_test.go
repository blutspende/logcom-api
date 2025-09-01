package logcom

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	logcomapi "github.com/blutspende/logcom-api"
	"github.com/google/uuid"
)

func TestSendAndNotifyAndLog(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "Authorization", "BearerToken")
		ctx = context.WithValue(ctx, "RequestID", uuid.NewString())

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(204)
		}))
		defer svr.Close()

		config := Configuration{
			ServiceName: "Unknown",
			LogComURL:   svr.URL,
		}
		Init(config)
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
	})
}
