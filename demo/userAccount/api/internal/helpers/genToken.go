package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenToken(name string, secret_key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Second * 60).Unix(),
	})
	return token.SignedString([]byte(secret_key))
}
