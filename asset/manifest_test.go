package asset

import (
	"os"
	"strings"
	"testing"
)

func TestProductionManifests(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	m := NewManifest("../public/assets/manifest.json")
	hashedAppJs := m.Path("app.js")

	if !strings.HasPrefix(hashedAppJs, "/assets/app.") {
		t.Errorf("'%s' is not start with 'app.'", hashedAppJs)
	}

	if !strings.HasSuffix(hashedAppJs, ".js") {
		t.Errorf("'%s' is not end with '.js'", hashedAppJs)
	}
	os.Unsetenv("APP_ENV")
}

func TestDevManifests(t *testing.T) {
	m := NewManifest("../public/assets/manifest.json")
	hashedAppJs := m.Path("app.js")

	if !strings.HasPrefix(hashedAppJs, "https://lvh.me:8080/assets/app.") {
		t.Errorf("'%s' is not start with 'app.'", hashedAppJs)
	}

	if !strings.HasSuffix(hashedAppJs, ".js") {
		t.Errorf("'%s' is not end with '.js'", hashedAppJs)
	}
}
