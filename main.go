package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	handler := helloHandler{}
	server := http.Server{
		Addr:    ":3000",
		Handler: &handler,
	}

	server.ListenAndServe()
}
