package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
	"strconv"
	"strings"
)

type EmployeeProvider struct {
	mapper *mappers.EmployeeMapper
}

func (e *EmployeeProvider) UpdateEmploy(Id, Surname, Name, Position string) (err error) {
	employee, err := e.mapper.GetEmployeeById(Id)

	employee.Surname = Surname
	employee.Name = Name
	employee.Position = Position

	err = e.mapper.EmployeeUpdate(&employee)

	return err
}

func (e *EmployeeProvider) NewEmployee(Surname, Name, Position string) (id int) {
	employee := &entities.Employee{Surname: Surname, Name: Name, Position: Position}

	id, err := e.mapper.EmployeeAdd(employee)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func (e *EmployeeProvider) GetIdDesignatedEmployee(employee string) (id string) {
	massSurName := strings.Split(employee, " ")

	surname := massSurName[0]
	name := massSurName[1]

	rowDesignatedEmployee, err := e.mapper.GetEmployeeBySurnameName(surname, name)
	if err != nil {
		fmt.Println(err)
	}

	id = strconv.Itoa(rowDesignatedEmployee.Id)
	return id
}
