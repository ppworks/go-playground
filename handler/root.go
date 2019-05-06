package handler

import (
	"html/template"
	"net/http"

	"github.com/ppworks/go-playground/asset"
	"github.com/ppworks/go-playground/handler/hook"
	"github.com/ppworks/go-playground/helper"
)

// RootHandlefunc is handler of "/"
func RootHandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	count := hook.CountUp("index")

	t := template.Must(template.New("").Funcs(helper.AppHelper()).ParseFiles(
		"templates/layouts/application.html",
		"templates/index.html",
	))

	t.ExecuteTemplate(w, "layout", struct {
		asset.ManifestFile
		BodyCSS     string
		AccessCount int64
	}{
		asset.ManifestFile{
			AppJS:  asset.Path("app.js"),
			AppCSS: asset.Path("app.css"),
		}, "text-center", count,
	})
}
