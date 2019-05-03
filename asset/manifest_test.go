package asset

import (
	"strings"
	"testing"
)

func TestManifests(t *testing.T) {
	m := NewManifest("../public/js/manifest.json")
	hashedAppJs := m.FileName("app.js")

	if !strings.HasPrefix(hashedAppJs, "app.") {
		t.Errorf("'%s' is not start with 'app.'", hashedAppJs)
	}

	if !strings.HasSuffix(hashedAppJs, ".js") {
		t.Errorf("'%s' is not end with '.js'", hashedAppJs)
	}
}
