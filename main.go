package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandlefunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/", rootHandlefunc)

	if os.Getenv("APP_ENV") == "production" {
		server.ListenAndServe()
	} else {
		server.ListenAndServeTLS("./ssl/lvh.me.crt", "./ssl/lvh.me.key")
	}
}
