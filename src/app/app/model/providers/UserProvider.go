package providers

import (
	"app/app"
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserProvider struct {
	mapper *mappers.UserMapper
}

//Добавление нового пользователя. Хэширование пароля.
func (p *UserProvider) NewUser(email, password string) (err error) {
	password = app.HashPassword([]byte(password))

	user := &entities.User{Email: email, Password: password}

	p.mapper = new(mappers.UserMapper)

	err = p.mapper.AddNewUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

//Сравнение данных пользователя, хранящихся в бд и введённых пользователем.
func (p *UserProvider) AutoriseUser(email, password string) bool {
	p.mapper = new(mappers.UserMapper)
	user, err := p.mapper.GetUserByEmail(email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true
	}
	return false
}
