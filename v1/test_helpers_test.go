package v1_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// APITestCase defines a single API contract test.
type APITestCase struct {
	Method  string
	Path    string
	Status  int
	Fixture string // optional JSON response fixture

	// Request validation hook
	ValidateRequest func(t *testing.T, r *http.Request)
}

// LoadsFixture loads a test fixture json.
func LoadFixture(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		t.Fatalf("failed to read fixture %s: %v", path, err)
	}

	return string(b)
}

// RunAPITest starts a test HTTP server and executes a single contract case.
func RunAPITest(t *testing.T, tc APITestCase) *httptest.Server {
	t.Helper()

	handler := http.NewServeMux()

	handler.HandleFunc(tc.Path, func(w http.ResponseWriter, r *http.Request) {
		// Method check
		require.Equal(t, tc.Method, r.Method)

		// Validate request if needed
		if tc.ValidateRequest != nil {
			tc.ValidateRequest(t, r)
		}

		// Write response
		w.WriteHeader(tc.Status)

		if tc.Fixture != "" {
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(tc.Fixture))
			require.NoError(t, err)
		}
	})

	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)

	return ts
}

// ExpectQuery asserts exact query parameters.
func ExpectQuery(t *testing.T, expected map[string]string) func(*http.Request) {
	t.Helper()

	return func(r *http.Request) {
		t.Helper()

		q := r.URL.Query()

		for key, want := range expected {
			got := q.Get(key)
			require.Equalf(t, want, got, "query param %q mismatch", key)
		}

		// optional strict mode: ensure no extra params
		if len(q) != len(expected) {
			t.Fatalf("unexpected query params: got=%v expected=%v", q, expected)
		}
	}
}

// ExpectJSONBody asserts request body matches expected struct.
func ExpectJSONBody[T any](expected T) func(*testing.T, *http.Request) {
	return func(t *testing.T, r *http.Request) {
		t.Helper()

		var got T
		err := json.NewDecoder(r.Body).Decode(&got)
		require.NoError(t, err)

		require.Equal(t, expected, got)
	}
}
