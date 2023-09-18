package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenJwt(issuerId int) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(issuerId),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
