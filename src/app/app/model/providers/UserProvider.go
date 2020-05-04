package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(email, password string) (err error) {
	password = HashPassword([]byte(password))

	user := &entities.User{Email: email, Password: password}

	err = mappers.AddNewUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func HashPassword(pass []byte) (password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, pass)

	password = string(hashedPassword)

	return password
}
