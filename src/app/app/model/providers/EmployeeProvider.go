package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
	"strconv"
	"strings"
)

func UpdateEmploy(Id, Surname, Name, Position string) (err error) {
	employee, err := mappers.GetEmployeeById(Id)

	employee.Surname = Surname
	employee.Name = Name
	employee.Position = Position

	err = mappers.EmployeeUpdate(&employee)

	return err
}

func NewEmployee(Surname, Name, Position string) (id int) {
	employee := &entities.Employee{Surname: Surname, Name: Name, Position: Position}

	id, err := mappers.EmployeeAdd(employee)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func GetIdDesignatedEmployee(employee string) (id string) {
	massSurName := strings.Split(employee, " ")

	surname := massSurName[0]
	name := massSurName[1]

	rowDesignatedEmployee, err := mappers.GetEmployeeBySurnameName(surname, name)
	if err != nil {
		fmt.Println(err)
	}

	id = strconv.Itoa(rowDesignatedEmployee.Id)
	return id
}
