package config

import (
	"net/http"

	"code.anexia.com/se/ks/go-anx-sdk/internal"
)

// Config is a collection of all configuration options for a new client.
// All fields are intentionally unexported and should be configured via the options pattern (see ClientOption).
type Config struct {
	apiKey  string
	baseUrl string
	client  *http.Client
}

func (c Config) CreateTransport() *internal.Transport {
	return internal.NewTransport(
		c.baseUrl,
		c.apiKey,
		nil,
	)
}

// ClientOption is a named type to improve code clarity and intent.
type ClientOption func(*Config)

// WithApiKey configures a new client with the given api key.
func WithApiKey(apiKey string) ClientOption {
	return func(config *Config) {
		config.apiKey = apiKey
	}
}

// WithBaseURL configures a new client with the given url as the base url.
func WithBaseURL(url string) ClientOption {
	return func(c *Config) {
		c.baseUrl = url
	}
}

// WithHttpClient configures a new client with the given http client.
func WithHttpClient(client *http.Client) ClientOption {
	return func(c *Config) {
		c.client = client
	}
}
