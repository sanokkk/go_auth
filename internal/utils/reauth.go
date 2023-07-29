package utils

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func Reauth(refresh string, secret string) (string, string, error) {
	claims := RefreshClaims{}
	token, err := jwt.ParseWithClaims(refresh, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("error while getting claims from refresh token: ", err)
		return "", "", err
	}
	if !token.Valid {
		log.Println("refresh token is not valid now")
		return "", "", errors.New("invalid refresh token")
	}
	login := claims.Login
	if login == nil {
		return "", "", errors.New("no credentials in claims")
	}
	generator := SH256JWT{}
	jwtToken, err := generator.generateJWT(secret, login)
	if err != nil {
		log.Println("error while getting new jwt: ", err)
		return "", "", err
	}
	newRefresh, _, err := generator.generateRefresh(secret, login)
	if err != nil {
		log.Println("error while getting new refresh: ", err)
		return "", "", err
	}
	return jwtToken, newRefresh, nil
}
