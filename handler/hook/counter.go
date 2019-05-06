package hook

import (
	"github.com/ppworks/go-playground/counter"
)

// CountUp specific page and return new counter value
func CountUp(pageName string) int64 {
	pageCounter := counter.NewPageCounter(pageName)
	pageCounter.Count()
	return pageCounter.Current
}
