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

func (p *EmployeeProvider) GetEmployees() (employees []entities.Employee, err error) {
	p.mapper = new(mappers.EmployeeMapper)

	employees, err = p.mapper.GetAllEmployees()
	if err != nil {
		fmt.Println(err)
		return employees, err
	}

	return employees, err
}

func (p *EmployeeProvider) UpdateEmploy(Id, Surname, Name, Position string) (err error) {
	p.mapper = new(mappers.EmployeeMapper)

	employee, err := p.mapper.GetEmployeeById(Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	employee.Surname = Surname
	employee.Name = Name
	employee.Position = Position

	err = p.mapper.EmployeeUpdate(&employee)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (p *EmployeeProvider) NewEmployee(Surname, Name, Position string) (id int, err error) {
	p.mapper = new(mappers.EmployeeMapper)

	employee := &entities.Employee{Surname: Surname, Name: Name, Position: Position}

	id, err = p.mapper.EmployeeAdd(employee)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, err
}

func (p *EmployeeProvider) GetIdDesignatedEmployee(employee string) (id string, err error) {
	p.mapper = new(mappers.EmployeeMapper)

	massSurName := strings.Split(employee, " ")

	surname := massSurName[0]
	name := massSurName[1]

	rowDesignatedEmployee, err := p.mapper.GetEmployeeBySurnameName(surname, name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	id = strconv.Itoa(rowDesignatedEmployee.Id)
	return id, err
}

func (p *EmployeeProvider) DelEmployee(id string) (err error) {
	p.mapper = new(mappers.EmployeeMapper)

	err = p.mapper.EmployeeDelete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}
