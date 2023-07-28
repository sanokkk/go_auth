package models

import (
	"github.com/google/uuid"
	"github.com/sanokkk/go_auth/internal/db/repo"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"full_name"`
	EMail        string    `json:"email"`
	NickName     string    `json:"nick_name"`
	Age          int       `json:"age"`
	PasswordHash string    `json:"password_hash"`
}

type UserRegister struct {
	FullName        string `json:"full_name"`
	EMail           string `json:"e_mail"`
	NickName        string `json:"nick_name"`
	Age             int    `json:"age"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type UserLogin struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

func ConvertToMyUser(model *repo.User) *User {
	return &User{
		ID:           model.ID,
		FullName:     model.FullName,
		EMail:        model.Email,
		NickName:     model.NickName,
		Age:          int(model.Age),
		PasswordHash: model.PasswordHash,
	}
}
