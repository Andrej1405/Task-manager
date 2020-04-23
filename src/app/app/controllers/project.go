package controllers

import (
	entities "app/app/model"
	"app/app/server"
	"fmt"

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

func (c ControllerProject) AddNewProject(Name, Date string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	project := entities.NewProject(Name, Date)
	fmt.Println(project)
	id, err := entities.ProjectAdd(project)
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(id)
}

func (c ControllerProject) UpdateProject(Id, Name, Date string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	project, err := entities.GetProjectById(Id)
	if err != nil {
		panic(err)
	}

	project.Name = Name
	project.Date = Date

	err = entities.ProjectUpdate(&project)
	if err != nil {
		panic(err)
	}
	return c.Render()
}

func (c ControllerProject) DeleteProject(Id string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	err = entities.ProjectDelete(Id)
	if err != nil {
		panic(err)
	}

	return c.Render()
}
