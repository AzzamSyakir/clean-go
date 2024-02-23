package middleware

import (
	"clean-go/cache"
	"clean-go/internal/gateway/responses"
	"fmt"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		// Membersihkan token dari string "Bearer "
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		if tokenString == "" {
			responses.ErrorResponse(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}
		// Cek token di Redis
		_, err := cache.GetCached(tokenString)
		if err == redis.Nil {
			http.Error(w, "Unauthorized: Token not found or expired", http.StatusUnauthorized)
			return
		} else if err != nil {
			fmt.Printf("Error accessing Redis: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Token valid, lanjutkan ke handler selanjutnya
		next.ServeHTTP(w, r)
	})
}
