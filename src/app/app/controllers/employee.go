package controllers

import (
	"github.com/revel/revel"
)

type ControllerProject struct {
	*revel.Controller
}
type Employee struct {
	Id       int
	Surname  string
	Name     string
	Position string
}

func (c ControllerProject) GetEmlpoyee() revel.Result {
	//employees := c.Params.Values
	employees := []Employee{
		{1, "Джон", "До", "Доктор"},
		{2, "Дориан", "Грей", "Водитель"},
	}
	return c.RenderJSON(employees)
}
