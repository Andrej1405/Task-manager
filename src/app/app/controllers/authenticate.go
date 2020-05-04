package controllers

import (
	"app/app"
	mappers "app/app/model/mappers"
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type Authenticate struct {
	*revel.Controller
}

func (c Authenticate) Sign() revel.Result {
	if app.GetSession(c.Session.ID()) {
		return c.Redirect((*App).Index)
	}
	return c.Render()
}

func (c Authenticate) Registration(Email, Password string) revel.Result {
	if Email == "" || Password == "" {
		message := "Данные некорректны. Добавить нового пользователя не удалось"
		fmt.Println(message)
		return c.RenderJSON(message)
	}

	email := Email
	password := Password

	existingUser, err := mappers.GetUserByEmail(email)
	if err == nil {
		fmt.Println(existingUser)
		message := "Пользователь с таким email уже зарегистрирован"
		return c.RenderJSON(message)
	}

	err = providers.NewUser(email, password)
	if err != nil {
		return c.RenderJSON(err)
	}

	return c.Render()
}

func (c Authenticate) Login(Email, Password string) revel.Result {
	if Email == "" || Password == "" {
		message := "Введены некорректные данные"
		fmt.Println(message)
		return c.RenderJSON(message)
	}

	email := Email
	password := Password

	boolean := mappers.AutoriseUser(email, password)

	if boolean == true {
		app.Add(c.Session.ID())
		return c.RenderJSON(true)
	}
	return c.RenderJSON(false)
}

func (c Authenticate) Logout() revel.Result {
	app.DeleteSession(c.Session.ID())
	return c.Render()
}
