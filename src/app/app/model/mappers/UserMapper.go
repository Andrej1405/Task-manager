package mappers

import (
	"app/app/model/entities"
	"app/app/model/providers"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(email string) (user entities.User, err error) {
	row := providers.GetUserByEmail(email)
	user = entities.User{}

	err = row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func NewUser(email, password string) *entities.User {
	return &entities.User{Email: email, Password: password}
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

func AutoriseUser(email, password string) bool {
	row := providers.GetUserByEmail(email)
	user := entities.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true
	}
	return false
}
