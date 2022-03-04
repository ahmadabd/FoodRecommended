package yaml

import (
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
)

func TestGetServerHost(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetServerHost() != "localhost" {
		t.Errorf("got, want %s", cfg.GetServerHost())
	}
}

func TestGetDatabaseHost(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetDatabaseHost() != "localhost" {
		t.Errorf("got, want %s", cfg.GetDatabaseHost())
	}
}

func TestGetServerPort(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetServerPort() != "5000" {
		t.Errorf("got, want %s", cfg.GetServerPort())
	}
}

func TestGetDatabasePort(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetDatabasePort() != "3306" {
		t.Errorf("got, want %s", cfg.GetDatabasePort())
	}
}

func TestGetDatabaseName(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetDatabaseName() != "FoodRecommendationTest" {
		t.Errorf("got, want %s", cfg.GetDatabaseName())
	}
}

func TestGetDatabaseUser(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetDatabaseUser() != "root" {
		t.Errorf("got, want %s", cfg.GetDatabaseUser())
	}
}

func TestGetDatabasePassword(t *testing.T) {
	path := "testdata/test.yml"

	cfg, _ := GetConfig(path)

	if cfg.GetDatabasePassword() != "1234" {
		t.Errorf("got, want %s", cfg.GetDatabasePassword())
	}
}

func TestGetConfig(t *testing.T) {
	senarios := []struct {
		name     string
		path     string
		expected configs.ConfigImp
		wantErr  bool
	}{
		{
			name:     "parse test.yml",
			path:     "testdata/config.yml",
			expected: &Config{},
			wantErr:  false,
		},
		{
			name:     "parse test.yaml",
			path:     "testdata/config.yaml",
			expected: &Config{},
			wantErr:  false,
		},
		{
			name:     "parse test.cfg",
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
}
