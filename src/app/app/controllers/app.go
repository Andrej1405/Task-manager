package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// if true {
	// 	return c.Redirect((*Authenticate).Login)
	// }
	return c.Render()
}
