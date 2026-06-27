package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey 1234567890")

	got, _ := GetAPIKey(headers)
	want := "1234567890"

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
