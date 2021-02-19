package metric

import (
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	errInvalidToken = errors.New("Invalid Token")
	errAccessDenied = errors.New("Access Denied")
)

// prometheusMiddleware implements mux.MiddlewareFunc.
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(AuthCount)
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
