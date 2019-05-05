package counter

import (
	"github.com/ppworks/go-playground/database"
)

// PageCounter ...
type PageCounter struct {
	Page    string
	Current int64
}

// NewPageCounter return PageCounter ref
func NewPageCounter(page string) *PageCounter {
	return &PageCounter{Page: page, Current: 0}
}

// Count increment page's counter
func (p *PageCounter) Count() {
	result, err := database.Redis().Incr(p.keyName()).Result()
	if err != nil {
		p.Current = 0
		err := database.Redis().Set(p.keyName(), 0, 0).Err()
		if err != nil {
			panic(err)
		}
	} else {
		p.Current = result
	}

	return
}

func (p *PageCounter) keyName() string {
	return "PageCounter:" + p.Page + ":counter"
}
