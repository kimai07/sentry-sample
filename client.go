package main

import (
  "log"
  "fmt"
  raven "github.com/getsentry/raven-go"
)

func main() {
  sentryClient := raven.DefaultClient
  err := sentryClient.SetDSN("<YOUR DSN>")
  if err != nil {
    log.Fatalf("failed to create sentry client: %v", err)
  }
  sentryClient.SetEnvironment("development")
  sentryClient.SetRelease("local_test")

  tags := map[string]string{}
  tags["environment"] = "development"
  tags["release"] = "local_test"
  tags["hoge"] = "fuga"

  err = fmt.Errorf("test error")
  res := sentryClient.CaptureError(err, tags)
  log.Println(res)
  res = sentryClient.CaptureErrorAndWait(err, tags)
  log.Println(res)
}

