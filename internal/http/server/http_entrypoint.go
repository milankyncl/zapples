package server

import (
	"net/http"
	"net/url"
)

func FileEntrypoint(filename string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = filename
		r2.URL.RawPath = filename
		h.ServeHTTP(w, r2)
	})
}
