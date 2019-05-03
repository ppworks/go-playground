package asset

import (
	"os"
	"strings"
	"testing"
)

func TestProductionManifests(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	m := NewManifest("../public/assets/manifest.json")

	cases := []struct {
		fileName     string
		resultPrefix string
		resultSuffix string
	}{
		{fileName: "app.js", resultPrefix: "/assets/app.", resultSuffix: ".js"},
		{fileName: "app.css", resultPrefix: "/assets/app.", resultSuffix: ".css"},
	}

	for _, c := range cases {
		result := m.Path(c.fileName)

		if !strings.HasPrefix(result, c.resultPrefix) {
			t.Errorf("'%s' is not start with 'app.'", result)
		}

		if !strings.HasSuffix(result, c.resultSuffix) {
			t.Errorf("'%s' is not end with '.js'", result)
		}
	}

	os.Unsetenv("APP_ENV")
}

func TestDevManifests(t *testing.T) {
	m := NewManifest("../public/assets/manifest.json")

	cases := []struct {
		fileName     string
		resultPrefix string
		resultSuffix string
	}{
		{fileName: "app.js", resultPrefix: "https://lvh.me:8080/assets/app.", resultSuffix: ".js"},
		{fileName: "app.css", resultPrefix: "https://lvh.me:8080/assets/app.", resultSuffix: ".css"},
	}

	for _, c := range cases {
		result := m.Path(c.fileName)

		if !strings.HasPrefix(result, c.resultPrefix) {
			t.Errorf("'%s' is not start with 'app.'", result)
		}

		if !strings.HasSuffix(result, c.resultSuffix) {
			t.Errorf("'%s' is not end with '.js'", result)
		}
	}
}
