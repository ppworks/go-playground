package main

import (
	"fmt"
	"net/http"
	"os"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	handler := helloHandler{}
	server := http.Server{
		Addr:    ":" + port,
		Handler: &handler,
	}

	server.ListenAndServe()
}
