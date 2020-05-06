package controllers

import (
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerEmployee struct {
	*revel.Controller
	provider *providers.EmployeeProvider
}

// Передача списка сотрудников на фронт.
func (c ControllerEmployee) GetEmployee() revel.Result {
	c.provider = new(providers.EmployeeProvider)

	employees, err := c.provider.GetEmployees()
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.RenderJSON(employees)
}

// Добавление нового сотрудника. Возвращает id добавленного сотрудника.
func (c ControllerEmployee) AddNewEmployee(Surname, Name, Position string) revel.Result {
	if Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Добавить сотрудника не удалось")
		return c.Render()
	}

	c.provider = new(providers.EmployeeProvider)

	id, err := c.provider.NewEmployee(Surname, Name, Position)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}
	return c.RenderJSON(id)
}

// Удаление сотрудника по его id.
func (c ControllerEmployee) DeleteEmployee(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id пустой. Удалить сотрудника не удалось")
		return c.Render()
	}

	c.provider = new(providers.EmployeeProvider)

	err := c.provider.DelEmployee(Id)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.Render()
}

// Обновление информации по существующему сотруднику.
func (c ControllerEmployee) UpdateEmployee(Id, Surname, Name, Position string) revel.Result {
	if Id == "" || Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Обновление сотрудника не удалось")
		return c.Render()
	}

	c.provider = new(providers.EmployeeProvider)

	err := c.provider.UpdateEmploy(Id, Surname, Name, Position)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.Render()
}
