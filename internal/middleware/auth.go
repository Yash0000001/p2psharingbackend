package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized user", err)
			return
		}
		tokenStr := cookie.Value
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			utils.SendError(w, http.StatusUnauthorized, "Invalid token", err)
		}

		userID, ok := claims["userID"].(string)

		if !ok {
			utils.SendError(w, http.StatusUnauthorized, "Invalid token payload", err)
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
