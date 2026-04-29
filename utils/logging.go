package utils

import (
	"log"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Next http.RoundTripper
}

func NewLoggingRoundTripper(next http.RoundTripper) *LoggingRoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}

	return &LoggingRoundTripper{
		Next: next,
	}
}

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
