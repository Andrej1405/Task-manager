package controllers

import (
	mappers "app/app/model/mappers"
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerEmployee struct {
	*revel.Controller
}

func (c ControllerEmployee) GetEmployee() revel.Result {
	employees, err := mappers.GetAllEmployees()
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.RenderJSON(employees)
}

func (c ControllerEmployee) AddNewEmployee(Surname, Name, Position string) revel.Result {
	if Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Добавить сотрудника не удалось")
		return c.Render()
	}

	id := providers.NewEmployee(Surname, Name, Position)

	return c.RenderJSON(id)
}

func (c ControllerEmployee) DeleteEmployee(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id пустой. Удалить сотрудника не удалось")
		return c.Render()
	}

	err := mappers.EmployeeDelete(Id)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}

func (c ControllerEmployee) UpdateEmployee(Id, Surname, Name, Position string) revel.Result {
	if Id == "" || Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Обновление сотрудника не удалось")
		return c.Render()
	}

	err := providers.UpdateEmploy(Id, Surname, Name, Position)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
