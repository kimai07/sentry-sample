package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "http://localhost:9000/", // ***YOUR_SENTRY_DSN***
		Environment:      "dev",
		Release:          "v0.0.1",
		Debug:            true,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	eventID := sentry.CaptureMessage("test message")
	fmt.Printf("event id: %v\n", eventID)

	event := &sentry.Event{
		Level:   sentry.LevelInfo,
		Message: "test hello",
		Request: &sentry.Request{
			URL:    "/test-hello",
			Method: "GET",
			Data:   "payload",
			Headers: map[string]string{
				"Accept-Encoding": "gzip",
				"Content-Length":  "7",
				"User-Agent":      "Go-http-client/1.1",
			},
		},
		Transaction: "GET /test-hello",
	}
	eventID = sentry.CaptureEvent(event)
	fmt.Printf("event id: %v\n", eventID)

	eventID = sentry.CaptureException(fmt.Errorf("test error"))
	fmt.Printf("event id: %v\n", eventID)

}
