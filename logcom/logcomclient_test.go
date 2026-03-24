package logcom

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	logcomapi "github.com/blutspende/logcom-api"
)

func TestSendConsoleLogWithModelOk(t *testing.T) {
	ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")

	dto := logcomapi.CreateConsoleLogRequestDTO{
		Level:   logcomapi.Debug,
		Message: "Test send console log with model",
	}
	err := sendConsoleLogWithModel(ctx, dto)
	if err != nil {
		t.Errorf("expected no error")
	}
}

func TestSendConsoleLogWithModelBadRequest(t *testing.T) {
	badRequest = true
	ctx := context.WithValue(context.Background(), "Authorization", "BearerToken")
	dto := logcomapi.CreateConsoleLogRequestDTO{
		Level:   logcomapi.Debug,
		Message: "Test send console log with model",
	}
	expectedResponse := strconv.Itoa(http.StatusBadRequest) + " " + http.StatusText(http.StatusBadRequest)

	err := sendConsoleLogWithModel(ctx, dto)
	if err == nil || err.Error() != expectedResponse {
		t.Errorf("expected result to be %s got %s", strconv.Itoa(http.StatusBadRequest)+http.StatusText(http.StatusBadRequest), err)
	}

	badRequest = false
}
