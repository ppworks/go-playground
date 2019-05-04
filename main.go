package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/ppworks/go-playground/asset"
)

var manifest *asset.Manifest

func init() {
	manifest = asset.NewManifest("public/assets/manifest.json")

	redisURL := os.Getenv("REDIS_URL")
	opt, _ := redis.ParseURL(redisURL)
	client := redis.NewClient(opt)

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

}

func rootHandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles(
		"templates/layouts/application.html",
		"templates/index.html",
	))
	t.ExecuteTemplate(w, "layout", struct {
		AppJS, AppCSS, BodyCSS string
	}{manifest.Path("app.js"), manifest.Path("app.css"), "text-center"})
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
