package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GetEncrypted(src string) string {
	passwd, _ := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
	return string(passwd)
}

func CheckPW(encodePW, inputPW string) string {
	err := bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(inputPW))
	if err != nil {
		return "PW Wrong!"
	} else {
		return "PW Right!"
	}
}
