package logging

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogRequestDuration(severity string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logLevel, present := os.LookupEnv("LOG_LEVEL")
			if !present {
				logLevel = "ERROR"
			}

			// Log only if the environment variable LOG_LEVEL matches the severity
			switch logLevel {
			case "INFO":
				InfoLogger.Printf("%s %s %v", r.Method, r.RequestURI, time.Since(start))
			case "WARNING":
				WarningLogger.Printf("%s %s %v", r.Method, r.RequestURI, time.Since(start))
			case "ERROR":
				ErrorLogger.Printf("%s %s %v", r.Method, r.RequestURI, time.Since(start))
			default:
				log.Printf("Unknown severity: %s", severity)
			}

			next.ServeHTTP(w, r)
		})
	}
}
