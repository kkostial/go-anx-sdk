package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type TransportError struct {
	StatusCode int
	Status     string
	Body       string
}

func (a *TransportError) Error() string {
	return fmt.Sprintf("transport error: StatusCode=%d, Status=%s, Body=%s", a.StatusCode, a.Status, a.Body)
}

type Transport struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func (t *Transport) buildRequestURL(endpoint string, params any) (string, error) {
	base, err := url.Parse(t.baseURL)
	if err != nil {
		return "", fmt.Errorf("parsing base url: %w", err)
	}

	rel, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("parsing endpoint url: %w", err)
	}

	base = base.ResolveReference(rel)

	q, err := query.Values(params)
	if err != nil {
		return "", fmt.Errorf("building query params: %w", err)
	}
	base.RawQuery = q.Encode()

	fullURL := base.String()
	return fullURL, nil
}

func (t *Transport) newRequest(ctx context.Context, method, endpoint string, body io.Reader, params any) (*http.Request, error) {
	fullURL, err := t.buildRequestURL(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("building request url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	if t.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", t.apiKey))
	}

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (t *Transport) do(req *http.Request, response any) error {
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("reading response body: %w", err)
		}

		return &TransportError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	if response != nil {
		err = json.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			return fmt.Errorf("decoding body: %w", err)
		}
	}

	return nil
}

func (t *Transport) doWithBody(ctx context.Context, method string, endpoint string, request any, response any, params any) error {
	var body io.Reader

	if request != nil {
		var reqBody bytes.Buffer

		err := json.NewEncoder(&reqBody).Encode(request)
		if err != nil {
			return fmt.Errorf("marshalling request: %w", err)
		}

		body = &reqBody
	}

	req, err := t.newRequest(ctx, method, endpoint, body, params)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	err = t.do(req, response)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	return nil
}

func (t *Transport) Get(ctx context.Context, endpoint string, response any, params any) error {
	return t.doWithBody(ctx, http.MethodGet, endpoint, nil, response, params)
}

func (t *Transport) Delete(ctx context.Context, endpoint string) error {
	return t.doWithBody(ctx, http.MethodDelete, endpoint, nil, nil, nil)
}

func (t *Transport) Post(ctx context.Context, endpoint string, request any, response any) error {
	return t.doWithBody(ctx, http.MethodPost, endpoint, request, response, nil)
}

func (t *Transport) Put(ctx context.Context, endpoint string, request any, response any) error {
	return t.doWithBody(ctx, http.MethodPut, endpoint, request, response, nil)
}

// NewTransport creates a new internal transport helper with the provided base url, api key and http client.
func NewTransport(baseURL string, apiKey string, client *http.Client) *Transport {
	if client == nil {
		client = http.DefaultClient
	}

	return &Transport{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  client,
	}
}
