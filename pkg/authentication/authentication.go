package authentication

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	claims := CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
func GetEmailByToken(token string) (string, error) {
	arr := strings.Split(token, " ")
	if len(arr) <= 1 {
		return "", errors.New("invalid token")
	}
	token = arr[1]

	claims := CustomClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}
