package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/ppworks/go-playground/asset"
)

func rootHandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	manifest := asset.NewManifest("public/assets/manifest.json")
	t := template.Must(template.ParseFiles("templates/layouts/application.html"))
	t.ExecuteTemplate(w, "layout", struct {
		AppJs, AppCss string
	}{manifest.Path("app.js"), manifest.Path("app.css")})
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
