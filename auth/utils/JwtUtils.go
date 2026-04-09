package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//TODO: make this be an enviroment variable
var key = []byte("my_super_secret_key_for_hmac_veoma_tajno_ne_pokazuj_nikome")

func GenerateToken(username string) (string, error){
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ValidateToken(token string) {
}
