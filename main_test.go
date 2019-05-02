package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandlefunc))
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if "Hello World" != string(body) {
		t.Errorf("actual body: %v", string(body))
	}
}

func TestNotFoundPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandlefunc))
	defer ts.Close()

	r, err := http.Get(ts.URL + "/hoge")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if 404 != r.StatusCode {
		t.Errorf("statusCode: 404 != %v", r.StatusCode)
	}
}
