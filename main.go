package main

import (
	"io"
	"net/http"
	"time"
)

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// simulate run time
	time.Sleep(time.Millisecond * 300)

	io.WriteString(w, "Foo!\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, "Bar!\n")
}

func main() {
	http.HandleFunc("/foo", wrap(foo))
	http.HandleFunc("/bar", wrap(bar))
	http.ListenAndServe(":9090", nil)
}
