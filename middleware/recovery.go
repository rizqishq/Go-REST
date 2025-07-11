package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				// Return a 500 error
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{
					Error:   "Internal Server Error",
					Message: "An unexpected error occurred",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}