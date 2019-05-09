package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ppworks/go-playground/handler"
)

func TestRootHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.RootHandleFunc))
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if 200 != r.StatusCode {
		t.Errorf("statusCode: 202 != %v", r.StatusCode)
	}
}

func TestNotFoundPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.RootHandleFunc))
	defer ts.Close()

	r, err := http.Get(ts.URL + "/hoge")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if 404 != r.StatusCode {
		t.Errorf("statusCode: 404 != %v", r.StatusCode)
	}
}
