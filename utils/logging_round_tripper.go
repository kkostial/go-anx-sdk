package utils

import (
	"log"
	"net/http"
	"time"
)

// LoggingRoundTripper is a helper http.RoundTripper implementation that logs requests and responses.
type LoggingRoundTripper struct {
	Next http.RoundTripper
}

// NewLoggingRoundTripper creates a new instance of a LoggingRoundTripper.
func NewLoggingRoundTripper(next http.RoundTripper) *LoggingRoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}

	return &LoggingRoundTripper{
		Next: next,
	}
}

// RoundTrip implements the http.RoundTripper interface for LoggingRoundTripper.
func (t *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	log.Printf("%s %s", req.Method, req.URL.String())

	resp, err := t.Next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %s", resp.Status, time.Since(start))
	return resp, nil
}
