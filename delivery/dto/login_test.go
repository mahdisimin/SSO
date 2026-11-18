package dto

import (
	"strings"
	"testing"
)

func TestLoginRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request LoginRequest
		wantErr bool
	}{
		{"valid", LoginRequest{"mahdi", "1234"}, false},
		{"empty username", LoginRequest{"", "1234"}, true},
		{"spaces only", LoginRequest{"   ", "1234"}, true},
		{"empty password", LoginRequest{"mahdi", ""}, true},
		{"password too long", LoginRequest{"mahdi", strings.Repeat("a", 73)}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.request.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("LoginRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
