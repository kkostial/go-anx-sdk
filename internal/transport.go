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
	baseUrl string
	apiKey  string
	client  *http.Client
}

func (t *Transport) BuildRequestUrl(endpoint string, params any) (string, error) {
	base, err := url.Parse(t.baseUrl)
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

	fullUrl := base.String()
	return fullUrl, nil
}

func (t *Transport) NewRequest(ctx context.Context, method, endpoint string, body io.Reader, params any) (*http.Request, error) {
	fullUrl, err := t.BuildRequestUrl(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("building request url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullUrl, body)
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

func (t *Transport) Do(req *http.Request, response any) error {
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

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

func (t *Transport) Get(ctx context.Context, endpoint string, response any, params any) error {
	return t.DoWithBody(ctx, http.MethodGet, endpoint, nil, response, params)
}

func (t *Transport) Delete(ctx context.Context, endpoint string) error {
	return t.DoWithBody(ctx, http.MethodDelete, endpoint, nil, nil, nil)
}

func (t *Transport) Post(ctx context.Context, endpoint string, request any, response any) error {
	return t.DoWithBody(ctx, http.MethodPost, endpoint, request, response, nil)
}

func (t *Transport) Put(ctx context.Context, endpoint string, request any, response any) error {
	return t.DoWithBody(ctx, http.MethodPut, endpoint, request, response, nil)
}

func (t *Transport) DoWithBody(ctx context.Context, method string, endpoint string, request any, response any, params any) error {
	var body io.Reader

	if request != nil {
		var reqBody bytes.Buffer

		err := json.NewEncoder(&reqBody).Encode(request)
		if err != nil {
			return fmt.Errorf("marshalling request: %w", err)
		}

		body = &reqBody
	}

	req, err := t.NewRequest(ctx, method, endpoint, body, params)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	err = t.Do(req, response)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	return nil
}

func NewTransport(baseUrl string, apiKey string, client *http.Client) *Transport {
	if client == nil {
		client = http.DefaultClient
	}

	return &Transport{
		baseUrl: baseUrl,
		apiKey:  apiKey,
		client:  client,
	}
}
