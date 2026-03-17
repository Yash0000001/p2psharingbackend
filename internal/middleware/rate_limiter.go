package middleware

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(time.Minute/5), 5)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})

}
