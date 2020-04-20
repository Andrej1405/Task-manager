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
		panic(err)
	}

	employees, err := entities.GetAllEmployees()
	if err != nil {
		panic(err)
	}

	return c.RenderJSON(employees)
}

func (c ControllerEmployee) AddNewEmployee(Surname, Name, Position string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	employee := entities.NewEmployee(Surname, Name, Position)

	err = entities.EmployeeAdd(employee)
	if err != nil {
		panic(err)
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	// err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	// if err != nil {
	// 	http.Error(rw, err.Error(), 400)
	// 	return
	// }
	return c.Render()
}

func (c ControllerEmployee) UpdateEmployee(Id, Surname, Name, Position string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	employee, err := entities.GetEmployeeById(Id)
	if err != nil {
		panic(err)
	}

	employee.Surname = Surname
	employee.Name = Name
	employee.Position = Position

	err = entities.EmployeeUpdate(&employee)
	if err != nil {
		panic(err)
	}
	return c.Render()
}

func (c ControllerEmployee) DeleteUmployee(Id, Surname, Name, Position string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(Id, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	// employee, err := entities.GetEmployeeById(Id)
	// if err != nil {
	// 	panic(err)
	// }

	err = entities.EmployeeDelete(Id)
	if err != nil {
		panic(err)
	}

	return c.Render()
}
