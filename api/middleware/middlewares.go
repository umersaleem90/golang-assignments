package middleware

import (
	"net/http"
	"todo-backend/api/response"
	"todo-backend/utilities"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func ValidateUser(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, err := utilities.VerifyToken(r)
        if err != nil {
            response.WriteError(w, http.StatusUnauthorized, err)
            return
        }
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
