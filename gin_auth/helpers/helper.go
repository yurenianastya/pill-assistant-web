package helpers

import (
	"awesomeProject/controller"
	"log"
	"strings"
)

func CheckUserPass(username, password string) bool {
	log.Println("checkUserPass", username, password)
	passCheck, err := controller.FindByField("username", username)
	if err != nil {
		return false
	}
	if passCheck == "" {
		return false
	}
	return true
}

func EmptyUserPassLogin(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func EmptyUserPassSignUp(username, password, email string) bool {
	return strings.Trim(username, " ") == "" ||
		strings.Trim(password, " ") == "" || strings.Trim(email, " ") == ""
}
