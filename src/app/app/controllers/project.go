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
	projects, err := mappers.GetAllProject()
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

	id := providers.NewProject(Name)

	return c.RenderJSON(id)
}

func (c ControllerProject) DeleteProject(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id проекта пустой. Удалить проект не удалось")
		return c.Render()
	}

	err := mappers.ProjectDelete(Id)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}

func (c ControllerProject) UpdateProject(Id, Name string) revel.Result {
	if Name == "" {
		fmt.Println("Название проекта пусто. Обновление проекта не удалось")
		return c.Render()
	}

	err := providers.UpdateProj(Id, Name)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
