package main

import (
	"net/http"
	"os"

	"github.com/ppworks/go-playground/handler"
)

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

	http.HandleFunc("/", handler.RootHandlefunc)

	if os.Getenv("APP_ENV") == "production" {
		server.ListenAndServe()
	} else {
		server.ListenAndServeTLS("./ssl/lvh.me.crt", "./ssl/lvh.me.key")
	}
}
