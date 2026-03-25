package logcom

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup
	badRequest = false
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if badRequest {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(204)
		}
	}))
	config := Configuration{
		ServiceName: "Test",
		LogComURL:   testServer.URL,
	}
	Init(config)
	time.Sleep(time.Second)
	// Run
	code := m.Run()
	// Cleanup
	testServer.Close()
	os.Exit(code)
}

var badRequest bool
