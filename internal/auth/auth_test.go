package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		want   string
		err    error
	}{
		"no header": {
			header: http.Header{},
			want:   "",
			err:    ErrNoAuthHeaderIncluded,
		},
		"no api key": {
			header: http.Header{
				"Authorization": {"Bearer 1234567890"},
			},
			want: "",
			err:  ErrMalformedAuthHeader,
		},
		"valid api key": {
			header: http.Header{
				"Authorization": {"ApiKey 1234567890"},
			},
			want: "1234567890",
			err:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if tt.err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("expected error %v, got %v", tt.err, err)
				}
			} else if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
