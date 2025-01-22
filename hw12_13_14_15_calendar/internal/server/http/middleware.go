package internalhttp

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		latency := time.Since(start)
		log.Printf(
			"IP: %s, Time: %s, Method: %s, Path: %s, HTTP Version: %s, Status Code: %d, Latency: %v, User-Agent: %s\n",
			r.RemoteAddr,
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			r.Proto,
			rw.status,
			latency,
			r.UserAgent(),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
