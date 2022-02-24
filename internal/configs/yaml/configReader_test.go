package yaml

import (
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
)

func TestGetConfig(t *testing.T) {
	senarios := []struct {
		name     string
		path     string
		expected configs.ConfigImp
		wantErr  bool
	}{
		{
			name:     "parse config.yml",
			path:     "testdata/config.yml",
			expected: &Config{},
			wantErr:  false,
		},
		{
			name:     "parse config.yaml",
			path:     "testdata/config.yaml",
			expected: &Config{},
			wantErr:  false,
		},
		{
			name:     "parse config.cfg",
			path:     "testdata/config.cfg",
			expected: &Config{},
			wantErr:  true,
		},
	}

	for _, senario := range senarios {
		t.Run(senario.name, func(t *testing.T) {
			_, err := GetConfig(senario.path)
			if err != nil && !senario.wantErr {
				t.Errorf("%s: got, want %v", senario.name, senario.expected)
			}
		})
	}

	for _, senario := range senarios {
		_, err := GetConfig(senario.path)
		if err != nil && !senario.wantErr {
			t.Errorf("%s: got, want %v", senario.name, senario.expected)
		}
	}
}
