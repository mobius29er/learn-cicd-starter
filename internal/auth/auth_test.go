package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}                            // Adjust this based on the context
	header.Add("Authorization", "ApiKey expected-key") // Example addition

	expectedKey := "expected-key"

	gotKey, _ := GetAPIKey(header) // Capture both return values, ignore the second if unneeded

	if gotKey != expectedKey {
		t.Errorf("GetAPIKey() = %v; want %v", gotKey, expectedKey)
	}
}
