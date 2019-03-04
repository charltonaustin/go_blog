package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappingResponseWriter struct {
	http.ResponseWriter
	Status int
}

// AccessLogger is a handler decorator which logs the method, path, duration, and
// status for every request
func AccessLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := r.Method
		p := r.URL.Path

		s := time.Now()
		wr := &wrappingResponseWriter{ResponseWriter: w, Status: http.StatusOK}
		h.ServeHTTP(wr, r)
		d := time.Since(s)

		log.Printf("%v %v %v %v", m, p, d, wr.Status)
	})
}
