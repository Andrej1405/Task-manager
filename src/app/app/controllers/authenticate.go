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
	mapper   *mappers.UserMapper
	provider *providers.UserProvider
}

// Если у пользователя отсутствует сессия, то происходит редирект на страницу входа в приложение.
func (c Authenticate) Sign() revel.Result {
	if app.GetSession(c.Session.ID()) {
		return c.Redirect((*App).Index)
	}
	return c.Render()
}

// Регистрация нового пользователя.
func (c Authenticate) Registration(Email, Password string) revel.Result {
	email := Email
	password := Password

	existingUser, err := c.mapper.GetUserByEmail(email)
	if err == nil {
		fmt.Println(existingUser)
		message := "Пользователь с таким email уже зарегистрирован"
		return c.RenderJSON(message)
	}

	c.provider = new(providers.UserProvider)
	err = c.provider.NewUser(email, password)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render()
}

// Вход пользователя на сайт. Проверка введённых email и пароля.
func (c Authenticate) Login(Email, Password string) revel.Result {
	email := Email
	password := Password

	c.provider = new(providers.UserProvider)
	boolean := c.provider.AutoriseUser(email, password)

	if boolean == true {
		app.Add(c.Session.ID())
		return c.RenderJSON(true)
	}
	return c.RenderJSON(false)
}

// Разлогин пользователя на сайте. Удаляется сессия пользователя.
func (c Authenticate) Logout() revel.Result {
	app.DeleteSession(c.Session.ID())
	return c.Render()
}
