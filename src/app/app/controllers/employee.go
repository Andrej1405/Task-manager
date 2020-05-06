package controllers

import (
	"app/app/model/providers"

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
		return c.RenderError(err)
	}

	return c.RenderJSON(employees)
}

// Добавление нового сотрудника. Возвращает id добавленного сотрудника.
func (c ControllerEmployee) AddNewEmployee(Surname, Name, Position string) revel.Result {
	c.provider = new(providers.EmployeeProvider)

	id, err := c.provider.NewEmployee(Surname, Name, Position)
	if err != nil {
		return c.RenderError(err)
	}
	return c.RenderJSON(id)
}

// Удаление сотрудника по его id.
func (c ControllerEmployee) DeleteEmployee(Id string) revel.Result {
	c.provider = new(providers.EmployeeProvider)

	err := c.provider.DelEmployee(Id)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render()
}

// Обновление информации по существующему сотруднику.
func (c ControllerEmployee) UpdateEmployee(Id, Surname, Name, Position string) revel.Result {
	c.provider = new(providers.EmployeeProvider)

	err := c.provider.UpdateEmploy(Id, Surname, Name, Position)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render()
}
