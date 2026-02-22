package middleware

import (
	"log"
	"net/http"
	"time"
)

const validAPIKey = "supersecret"

// Логирует каждый запрос: время, метод, путь
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] method=%s endpoint=%s duration=%s",
			start.Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
		next.ServeHTTP(w, r)
	})
}

// Проверяет X-API-KEY header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key == "" || key != validAPIKey {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "unauthorized: invalid or missing X-API-KEY"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
