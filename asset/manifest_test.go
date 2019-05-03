package asset

import (
	"strings"
	"testing"
)

func TestManifests(t *testing.T) {
	hashedAppJs := Path("app.js")

	if !strings.HasPrefix(hashedAppJs, "app.") {
		t.Errorf("'%s' is not start with 'app.'", hashedAppJs)
	}

	if !strings.HasSuffix(hashedAppJs, ".js") {
		t.Errorf("'%s' is not end with '.js'", hashedAppJs)
	}
}
