package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sanokkk/go_auth/internal/config"
	"github.com/sanokkk/go_auth/internal/db/repo"
	"github.com/sanokkk/go_auth/internal/models"
	"github.com/sanokkk/go_auth/internal/utils"
	"github.com/sanokkk/go_auth/internal/validation"
)

func (apiCfg *ApiConfig) handlreCreateUser(w http.ResponseWriter, r *http.Request) {
	params := models.UserRegister{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println("error while getting from json: ", err)
		respondWithError(w, 400, struct {
			errMessage string `json:"err_message"`
		}{errMessage: err.Error()})
		return
	}
	valid, errors := validation.IsRegisterUserValid(&params)
	if !valid {
		log.Printf("error while validating: %v", errors)
		respondWithError(w, 400, struct {
			errorsList []string `json:"errors_list"`
		}{errorsList: errors})
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), repo.CreateUserParams{
		ID:           uuid.New(),
		FullName:     params.FullName,
		Email:        params.EMail,
		NickName:     params.NickName,
		Age:          int16(params.Age),
		PasswordHash: HashPassword(Sha256Hash{}, params.Password),
	})
	if err != nil {
		log.Println("error while getting creating user in db: ", err)
		respondWithError(w, 400, struct {
			errMessage string `json:"err_message"`
		}{errMessage: err.Error()})
		return
	}

	respondWithJSON(w, 201, models.ConvertToMyUser(&user))

}

func (apiCfg *ApiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	params := models.UserLogin{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println("error while getting login model from request: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}
	secretKeyStr, err := config.GetKey()
	if err != nil {
		log.Println("error while getting security key: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), repo.GetUserParams{
		NickName:     params.NickName,
		PasswordHash: HashPassword(Sha256Hash{}, params.Password),
	})
	if err != nil {
		log.Println("error while getting user from db: ", err.Error())
		respondWithError(w, 400, err.Error())
		return
	}
	JWTToken, err := utils.GenerateJWT(&utils.SH256JWT{}, secretKeyStr, models.ConvertToMyUser(&user))
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	refreshToken, expires, err := utils.GenerateRefresh(&utils.SH256JWT{}, secretKeyStr, models.ConvertToMyUser(&user))

	if err != nil {
		log.Println("error while creating refreshToken")
		respondWithError(w, 400, err.Error())
		return
	}
	_, err = apiCfg.DB.CraeateTokens(r.Context(), repo.CraeateTokensParams{
		ExpiresAt:    expires,
		RefreshToken: refreshToken,
	})
	if err != nil {
		log.Println("error while adding token to db")
		respondWithError(w, 400, err.Error())
		return
	}
	respondWithJSON(w, 200, utils.LoginResponse{JWTToken: JWTToken, RefreshToken: refreshToken})
}

func (apiCfg *ApiConfig) handleWelcome(w http.ResponseWriter, r *http.Request, user repo.User) {
	respondWithJSON(w, 200, user)
}

func (apiCfg *ApiConfig) handleReauth(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		RefreshToken string `json:"refresh_token"`
	}
	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println("error while getting params from body: ", err.Error())
		respondWithError(w, 400, err.Error())
		return
	}
	secretKeyStr, err := config.GetKey()
	if err != nil {
		log.Println("error while getting security key: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}
	newToken, newRefresh, err := utils.Reauth(params.RefreshToken, secretKeyStr)
	if err != nil {
		log.Println("error while getting new tokens: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}
	if err = apiCfg.DB.DeleteToken(r.Context(), params.RefreshToken); err != nil {
		log.Println("error while removing old token: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}
	if _, err = apiCfg.DB.CraeateTokens(r.Context(), repo.CraeateTokensParams{
		ExpiresAt:    time.Now().Add(time.Hour * 6),
		RefreshToken: newRefresh,
	}); err != nil {
		log.Println("error while adding new refresh token: ", err.Error())
		respondWithError(w, 500, err.Error())
		return
	}
	respondWithJSON(w, 200, utils.LoginResponse{JWTToken: newToken, RefreshToken: newRefresh})

}
