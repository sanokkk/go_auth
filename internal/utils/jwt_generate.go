package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sanokkk/go_auth/internal/models"
)

type JWTGenerator interface {
	generateJWT(secretKey string, login *models.User) (string, error)
}

func GenerateJWT(implementation JWTGenerator, secretKey string, login *models.User) (string, error) {
	return implementation.generateJWT(secretKey, login)
}

type SH256JWT struct{}

type Claims struct {
	NickName string    `json:"nick_name"`
	Id       uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func (generator *SH256JWT) generateJWT(secretKey string, login *models.User) (string, error) {
	byteKey := []byte(secretKey)
	expirationTime := time.Now().Add(time.Minute * 1)
	claims := &Claims{
		NickName: login.NickName,
		Id:       login.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteKey)
	if err != nil {
		log.Fatal("error while creating jwt signed with secruty key")
		return "", nil
	}
	return tokenString, nil

}
