package v1_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "code.anexia.com/se/ks/go-anx-sdk/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code.anexia.com/se/ks/go-anx-sdk/internal"
)

func TestLocationsClient_List(t *testing.T) {
	// arrange
	require := require.New(t)
	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// API surface validation (this is the important part)
		assert.Equal(http.MethodGet, r.Method)
		assert.Equal("/api/core/v1/location.json", r.URL.Path)

		assert.Equal("2", r.URL.Query().Get("page"))
		assert.Equal("10", r.URL.Query().Get("limit"))
		assert.Equal("vienna", r.URL.Query().Get("search"))

		// mock API response (IMPORTANT: match real API shape)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"data": []map[string]any{
					{
						"identifier": "loc-1",
						"code":       "VIE",
						"name":       "Vienna",
					},
				},
			},
		})
	}))
	defer ts.Close()

	tr := internal.NewTransport(ts.URL, "token", ts.Client())
	client := v1.NewLocationsClient(tr)

	params := v1.LocationListParams{
		Page:   2,
		Limit:  10,
		Search: "vienna",
	}

	// act
	resp, err := client.List(context.Background(), params)

	// assert
	require.NoError(err)

	require.Len(resp.Data, 1)
	assert.Equal("VIE", resp.Data[0].Code)
	assert.Equal("Vienna", resp.Data[0].Name)
}
