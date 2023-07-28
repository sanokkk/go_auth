package internal

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sanokkk/go_auth/internal/config"
	"github.com/sanokkk/go_auth/internal/db/repo"
	"github.com/sanokkk/go_auth/internal/utils"
)

type httpHandler func(w http.ResponseWriter, r *http.Request, user repo.User)

func (apiCfg *ApiConfig) handleAuth(handler httpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenWithBearer := r.Header.Get("Authorization")
		if tokenWithBearer == "" {
			respondWithError(w, 401, "no token")
			return
		}

		arr := strings.Split(tokenWithBearer, " ")
		if arr[0] != "Bearer" {
			respondWithError(w, 401, "no token")
			return
		}
		tokenStr := arr[1]

		myClaims := utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, &myClaims, func(token *jwt.Token) (interface{}, error) {
			key, _ := config.GetKey()
			return []byte(key), nil
		})
		log.Println("token valid: ", token.Valid)
		if err != nil {
			respondWithError(w, 401, "error while parcing token")
			return
		}
		user, err := apiCfg.DB.GetUserById(r.Context(), myClaims.Id)
		if err != nil {
			respondWithError(w, 401, "error while parcing id from claims")
			return
		}
		handler(w, r, user)

	}
}
