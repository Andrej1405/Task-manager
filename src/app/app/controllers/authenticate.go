package controllers

import (
	"github.com/revel/revel"
)

type Authenticate struct {
	*revel.Controller
}

func (c Authenticate) Login() revel.Result {
	return c.Render()
}
