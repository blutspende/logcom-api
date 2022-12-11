package logcom

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	logcomapi "github.com/DRK-Blutspende-BaWueHe/logcom-api"
)

func TestSendConsoleLogWithModel(t *testing.T) {
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
		dto := logcomapi.CreateConsoleLogRequestDTO{
			Level:   logcomapi.Debug,
			Message: "Test send console log with model",
		}
		err := sendConsoleLogWithModel(ctx, dto)
		if err != nil {
			t.Errorf("expected no error")
		}
	})

	t.Run("Bad request", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(400)
		}))
		defer svr.Close()

		Init(Configuration{
			ServiceName: "Unknown",
			LogComURL:   svr.URL,
		})

		ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")
		dto := logcomapi.CreateConsoleLogRequestDTO{
			Level:   logcomapi.Debug,
			Message: "Test send console log with model",
		}

		expectedResponse := strconv.Itoa(http.StatusBadRequest) + http.StatusText(http.StatusBadRequest)

		err := sendConsoleLogWithModel(ctx, dto)
		if err != nil && err.Error() != expectedResponse {
			return
		}

		t.Errorf("expected result to be %s got %s", strconv.Itoa(http.StatusBadRequest)+http.StatusText(http.StatusBadRequest), err)
	})
}
