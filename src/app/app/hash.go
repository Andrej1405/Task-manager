package app

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass []byte) (password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, pass)

	password = string(hashedPassword)

	return password
}
