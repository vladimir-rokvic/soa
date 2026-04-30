package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//TODO: make this be an enviroment variable
var key = []byte("my_super_secret_key_for_hmac_veoma_tajno_ne_pokazuj_nikome")

func GenerateToken(id uuid.UUID) (string, error){
	claims := jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ValidateToken(token_string string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		fmt.Println("Error validating token")
		fmt.Println(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims, nil
	}

	return nil, errors.New("Error getting claims")
}

//Middleware koji se ubacuje tokom obrade protected subroutova
type contextKey string
//pokusao sam i sa obicnim stringom ali bi trebalo da je bolje ovako
const ClaimsKey contextKey = "claims"

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			auth_header := req.Header.Get("Authorization")
			if auth_header == "" {
				fmt.Println("Error getting auth headear")
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			//ako nema Bearer pucaj u nogu
			parts := strings.Split(auth_header, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				fmt.Println("Error getting splitting auth header")
				fmt.Println(auth_header)
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			token := parts[1] 

			claims, err := ValidateToken(token)
			if err != nil {
				fmt.Println("Invalid token")
				fmt.Println(err)
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(req.Context(), ClaimsKey, claims)
			next.ServeHTTP(writer, req.WithContext(ctx))
		})
}

func GetClaims(req *http.Request) (jwt.MapClaims, bool) {
	claims, ok := req.Context().Value(ClaimsKey).(jwt.MapClaims)
	return claims, ok
}
