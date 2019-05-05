package formatter

import (
	"math"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	cases := []struct {
		Number    int64
		Formatted string
	}{
		{Number: 0, Formatted: "0"},
		{Number: 1, Formatted: "1"},
		{Number: -1, Formatted: "-1"},
		{Number: 12, Formatted: "12"},
		{Number: -12, Formatted: "-12"},
		{Number: 123, Formatted: "123"},
		{Number: -123, Formatted: "-123"},
		{Number: 1234, Formatted: "1,234"},
		{Number: -1234, Formatted: "-1,234"},
		{Number: 1234567, Formatted: "1,234,567"},
		{Number: -1234567, Formatted: "-1,234,567"},
		{Number: 1234567890, Formatted: "1,234,567,890"},
		{Number: -1234567890, Formatted: "-1,234,567,890"},
		{Number: math.MaxInt64, Formatted: "9,223,372,036,854,775,807"},
		{Number: math.MinInt64, Formatted: "-9,223,372,036,854,775,808"},
	}

	for _, c := range cases {
		result := NumberFormat(c.Number)

		if result != c.Formatted {
			t.Errorf("'%s' does't match '%s'", result, c.Formatted)
		}
	}
}
