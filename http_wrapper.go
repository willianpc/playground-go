package main

import (
	"net/http"
	"time"
)

func wrap(fn http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ins := newInstrumenter("http", func(sp span) {
			time.Sleep(time.Millisecond * 2000)

			sp.tags["path"] = r.URL.Path
			sp.tags["host"] = r.Host
			sp.tags["method"] = r.Method
		})

		// Impossible overhead to avoid, because it is sent in the request.
		// Same for every time we need span correllation.
		w.Header().Add("FOO", "BAR")

		ins.begin()
		fn(w, r)
		go ins.end()
	}
}
