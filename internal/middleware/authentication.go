package middleware

import (
	"fmt"
	"memo/internal/util/message"
	"memo/internal/util/response"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CookieからJWTを取得
		cookie, err := r.Cookie("jwt")

		fmt.Println(cookie)
		if err != nil || cookie == nil {
			response.Error(w, http.StatusUnauthorized, message.ErrUnauthorized)
			return
		}
		// token取得
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			response.Error(w, http.StatusUnauthorized, message.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
