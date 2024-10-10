package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name:       "Valid API key header",
			headers:    http.Header{"Authorization": {"ApiKey validapikey"}},
			wantAPIKey: "validapikey",
			wantErr:    nil,
		},
		{
			name:       "No Authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// gotAPIKey, gotErr := GetAPIKey(tt.headers)

			// Check for correct API key
			// if gotAPIKey != tt.wantAPIKey {
			// 	t.Errorf("GetAPIKey() = %v, want %v", gotAPIKey, tt.wantAPIKey)
			// }
			//
			// // Check for correct error
			// if gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
			// 	t.Errorf("GetAPIKey() error = %v, wantErr %v", gotErr, tt.wantErr)
			// }
			//
			// // Handle case where one is nil and the other is not
			// if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) {
			// 	t.Errorf("GetAPIKey() error = %v, wantErr %v", gotErr, tt.wantErr)
			// }
		})
	}
}
