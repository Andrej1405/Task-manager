package controllers

import (
	mappers "app/app/model/mappers"
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerProject struct {
	*revel.Controller
}

func (c ControllerProject) GetProject() revel.Result {
	projects, err := mappers.GetAllProjects()
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

	project := mappers.NewProject(Name)

	id, err := mappers.ProjectAdd(project)
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

	project, err := mappers.GetProjectById(Id)
	if err != nil {
		fmt.Println(err)
	}

	project.Name = Name

	err = providers.ProjectRowUpdate(&project)
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

	err := providers.ProjectRowDelete(Id)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
