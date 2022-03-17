package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_homeHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(homeHandler))
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {}()

	if !strings.Contains(string(result), "container") {
		t.Logf("%v", result)
		t.Fatal("not in container")
	}
}
