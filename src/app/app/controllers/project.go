package controllers

import (
	entities "app/app/model"
	"app/app/server"

	"github.com/revel/revel"
)

type ControllerProject struct {
	*revel.Controller
}

func (c ControllerProject) GetProject() revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	projects, err := entities.GetAllProjects()
	if err != nil {
		panic(err)
	}

	return c.RenderJSON(projects)
}

//employees := c.Params.Values
