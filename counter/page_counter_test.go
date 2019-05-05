package counter

import (
	"testing"
)

func TestCount(t *testing.T) {
	counter := NewPageCounter("index")

	before := counter.Current
	counter.Count()
	after := counter.Current

	if after <= before {
		t.Error("failed to increment counter")
	}
}
