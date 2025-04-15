package auth

import (
	"net/http"
	"os"

	"go-basic-poc/internal/utils"

	"go.uber.org/zap"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.Logger.Error("Authorization header missing")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		secretKey := os.Getenv("CLIENT_SECRET")
		_, err := ValidateJWT(tokenString, secretKey)
		if err != nil {
			utils.Logger.Error("Invalid token", zap.Error(err)) // Corrected line
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
