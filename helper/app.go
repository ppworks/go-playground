package helper

import (
	"html/template"

	"github.com/ppworks/go-playground/formatter"
)

// AppHelper for common template Funcs
func AppHelper() template.FuncMap {
	return template.FuncMap{
		"numberFormat": formatter.NumberFormat,
	}
}
