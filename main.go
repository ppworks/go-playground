package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/ppworks/go-playground/asset"
	"github.com/ppworks/go-playground/counter"
	"github.com/ppworks/go-playground/formatter"
)

var manifest *asset.Manifest

func init() {
	manifest = asset.NewManifest("public/assets/manifest.json")
}

func rootHandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	pageCounter := counter.NewPageCounter("index")
	pageCounter.Count()
	count := pageCounter.Current

	t := template.Must(template.New("").Funcs(template.FuncMap{
		"numberFormat": formatter.NumberFormat,
	}).ParseFiles(
		"templates/layouts/application.html",
		"templates/index.html",
	))

	t.ExecuteTemplate(w, "layout", struct {
		AppJS, AppCSS, BodyCSS string
		AccessCount            int64
	}{manifest.Path("app.js"), manifest.Path("app.css"), "text-center", count})
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := http.Server{
		Addr: ":" + port,
	}

	staticFiles := http.FileServer(http.Dir("public"))
	http.Handle("/assets/", staticFiles)

	http.HandleFunc("/", rootHandlefunc)

	if os.Getenv("APP_ENV") == "production" {
		server.ListenAndServe()
	} else {
		server.ListenAndServeTLS("./ssl/lvh.me.crt", "./ssl/lvh.me.key")
	}
}
