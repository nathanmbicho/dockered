package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	myOs, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := "inside"

	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = "outside"
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, "Hello %v!\n", r.UserAgent())
	_, _ = fmt.Fprintf(w, "I'm running on %v/%v.\n", myOs, myArch)
	_, _ = fmt.Fprintf(w, "I'm running %v of a container.\n", inContainer)
}

func main() {
	http.HandleFunc("/", homeHandler)
	err := http.ListenAndServe(":3800", nil)
	if err != nil {
		log.Fatal(err)
	}
}
