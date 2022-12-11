package logcom

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	logcomapi "github.com/DRK-Blutspende-BaWueHe/logcom-api"
)

func TestSendAndLog(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "Authorization", "BearerToken")

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(204)
		}))
		defer svr.Close()

		config := Configuration{
			ServiceName: "Unknown",
			LogComURL:   svr.URL,
		}
		Init(config)
		err := Notify().
			Roles("test_role").
			Message("Test notification").
			WithContext(ctx).
			AndLog(logcomapi.Debug, "Debug log").
			Send()
		if err != nil {
			t.Errorf("expected no error")
		}
	})
}
