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
		fmt.Println(err)
	}

	projects, err := entities.GetAllProjects()
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(projects)
}

func (c ControllerProject) AddNewProject(Name string) revel.Result {
	if Name == "" {
		fmt.Println("Название проекта пусто. Добавить проект не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	project := entities.NewProject(Name)

	id, err := entities.ProjectAdd(project)
	if err != nil {
		fmt.Println(err)
	}
	return c.RenderJSON(id)
}

func (c ControllerProject) UpdateProject(Id, Name string) revel.Result {
	if Name == "" {
		fmt.Println("Название проекта пусто. Обновление проекта не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	project, err := entities.GetProjectById(Id)
	if err != nil {
		fmt.Println(err)
	}

	project.Name = Name

	err = entities.ProjectUpdate(&project)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render()
}

func (c ControllerProject) DeleteProject(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id проекта пустой. Удалить проект не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	err = entities.ProjectDelete(Id)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
