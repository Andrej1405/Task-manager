package controllers

import (
	"app/app"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if !app.GetSessionById(c.Session.ID()) {
		return c.Redirect((*Authenticate).Sign)
	}
	return c.Render()
}
