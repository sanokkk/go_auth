package models

import "github.com/google/uuid"

type User struct {
	ID              uuid.UUID `json:"id"`
	FullName        string    `json:"full_name"`
	EMail           string    `json:"e_mail"`
	NickName        string    `json:"nick_name"`
	Age             int       `json:"age"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
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
