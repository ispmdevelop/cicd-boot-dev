package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type TestCase struct {
		input    http.Header
		expected string
	}

	testCases := []TestCase{
		{http.Header{"Authorization": {"ApiKey token123"}}, "token123"},
		{http.Header{"Authorization": {"ApiKey "}}, ""},
	}

	for _, tc := range testCases {
		got, err := GetAPIKey(tc.input)
		if err != nil && tc.expected != "" {
			t.Errorf("GetAPIKey(%v) returned error: %v, want no error", tc.input, err)
		}
		if got != tc.expected {
			t.Errorf("GetAPIKey(%v) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}
