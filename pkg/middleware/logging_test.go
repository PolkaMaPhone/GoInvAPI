package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	tests := []struct {
		name     string
		severity string
		envVar   string
		method   string
		url      string
	}{
		{
			name:     "WithINFOSeverity",
			severity: "INFO",
			envVar:   "LOG_LEVEL",
			method:   http.MethodGet,
			url:      "/test/info",
		},
		{
			name:     "WithWARNINGSeverity",
			severity: "WARNING",
			envVar:   "LOG_LEVEL",
			method:   http.MethodPost,
			url:      "/test/warning",
		},
		{
			name:     "WithERRORSeverity",
			severity: "ERROR",
			envVar:   "LOG_LEVEL",
			method:   http.MethodDelete,
			url:      "/test/error",
		},
		{
			name:     "WithUnknownSeverity",
			severity: "UNKNOWN",
			envVar:   "LOG_LEVEL",
			method:   http.MethodGet,
			url:      "/test/unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			// Set env variable
			err = os.Setenv(test.envVar, test.severity)
			if err != nil {
				return
			}
			defer func(key string) {
				err := os.Unsetenv(key)
				if err != nil {
					return
				}
			}(test.envVar)

			handler := LoggingMiddleware(test.severity)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

			handler.ServeHTTP(rr, req)
		})
	}
}
