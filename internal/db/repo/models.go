// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repo

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	FullName     string
	Email        string
	NickName     string
	Age          int16
	PasswordHash string
}
