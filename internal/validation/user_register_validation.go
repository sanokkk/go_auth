package validation

import (
	"net/mail"
	"regexp"

	"github.com/sanokkk/go_auth/internal/models"
)

func validMail(email string) (bool, string) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, "EMail is incorrect"
	}
	return true, ""
}

func validFullName(fullname string) (bool, string) {
	pattern := "/[A-ZА-Я][a-zа-я]{2,30} [A-ZА-Я][a-zа-я]{1,30} [A-ZА-Я][a-zа-я]{2,30}"
	res, err := regexp.MatchString(pattern, fullname)
	if err != nil {
		return false, "FullName is incorrect"
	}
	if res == true {
		return true, ""
	}
	return false, "FullName is incorrect"
}

func validNick(nick string) (bool, string) {
	pattern := "^[A-Za-z][A-Za-z0-9_]{5,29}$"
	res, err := regexp.MatchString(pattern, nick)
	if err != nil {
		return false, "Nickname is incorrect"
	}
	if res == true {
		return true, ""
	}
	return false, "Nickname is incorrect"
}

func validAge(age int) (bool, string) {
	if age <= 5 || age >= 100 {
		return false, "incorrect age"
	}
	return true, ""
}

func isPasswordConfirmed(pass, confPass string) (bool, string) {
	if pass != confPass {
		return false, "password is not confirmed"
	}
	return true, ""
}

func IsRegisterUserValid(register *models.UserRegister) (bool, []string) {
	result := true
	errorsList := []string{}
	if res, err := validFullName(register.FullName); !res {
		result = false
		errorsList = append(errorsList, err)
	}
	if res, err := validNick(register.NickName); !res {
		result = false
		errorsList = append(errorsList, err)
	}
	if res, err := validAge(register.Age); !res {
		result = false
		errorsList = append(errorsList, err)
	}
	if res, err := isPasswordConfirmed(register.Password, register.PasswordConfirm); !res {
		result = false
		errorsList = append(errorsList, err)
	}
	return result, errorsList
}
