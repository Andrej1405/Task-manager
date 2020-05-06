package controllers

import (
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerProject struct {
	*revel.Controller
	provider *providers.ProjectProvider
}

// Возвращает перечень проектов на фронт.
func (c ControllerProject) GetProject() revel.Result {
	c.provider = new(providers.ProjectProvider)

	projects, err := c.provider.GetAllProjects()
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}
	return c.RenderJSON(projects)
}

// Добавление нового проекта, возвращает id проекта на фронт.
func (c ControllerProject) AddNewProject(Name string) revel.Result {
	if Name == "" {
		fmt.Println("Название проекта пусто. Добавить проект не удалось")
		return c.Render()
	}
	c.provider = new(providers.ProjectProvider)

	id, err := c.provider.NewProject(Name)
	if err != nil {
		fmt.Println(err)
		return c.Render(err)
	}

	return c.RenderJSON(id)
}

// Удаление проекта по его id.
func (c ControllerProject) DeleteProject(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id проекта пустой. Удалить проект не удалось")
		return c.Render()
	}
	c.provider = new(providers.ProjectProvider)

	err := c.provider.DeleteProjectById(Id)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.Render()
}

// Обновление информации существующего проекта.
func (c ControllerProject) UpdateProject(Id, Name string) revel.Result {
	if Name == "" {
		fmt.Println("Название проекта пусто. Обновление проекта не удалось")
		return c.Render()
	}

	c.provider = new(providers.ProjectProvider)

	err := c.provider.UpdateProj(Id, Name)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.Render()
}
