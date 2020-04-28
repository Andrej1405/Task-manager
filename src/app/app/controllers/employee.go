package controllers

import (
	entities "app/app/model"
	"app/app/server"
	"fmt"

	"github.com/revel/revel"
)

type ControllerEmployee struct {
	*revel.Controller
}

func (c ControllerEmployee) GetEmployee() revel.Result {
	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	employees, err := entities.GetAllEmployees()
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(employees)
}

func (c ControllerEmployee) AddNewEmployee(Surname, Name, Position string) revel.Result {
	if Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Добавить сотрудника не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	employee := entities.NewEmployee(Surname, Name, Position)

	id, err := entities.EmployeeAdd(employee)
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(id)
}

func (c ControllerEmployee) UpdateEmployee(Id, Surname, Name, Position string) revel.Result {
	if Id == "" || Surname == "" || Name == "" || Position == "" {
		fmt.Println("Данные некорректны. Обновление сотрудника не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	employee, err := entities.GetEmployeeById(Id)
	if err != nil {
		fmt.Println(err)
	}

	employee.Surname = Surname
	employee.Name = Name
	employee.Position = Position

	err = entities.EmployeeUpdate(&employee)

	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}

func (c ControllerEmployee) DeleteEmployee(Id string) revel.Result {
	if Id == "" {
		fmt.Println("Id пустой. Удалить сотрудника не удалось")
		return c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	err = entities.EmployeeDelete(Id)

	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
