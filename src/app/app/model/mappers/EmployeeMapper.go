package mappers

import (
	"app/app/model/entities"
	"app/app/model/providers"
	"fmt"
	"strconv"
	"strings"
)

func GetAllEmployees() (employees []entities.Employee, err error) {
	employee := entities.Employee{}
	rows := providers.GetAllRowsEmployee()

	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
		if err != nil {
			fmt.Println(err)
		}
		employees = append(employees, employee)
	}

	return employees, err
}

func GetEmployeeById(employeeId string) (employee entities.Employee, err error) {
	employee = entities.Employee{}
	row := providers.GetRowEmployeeById(employeeId)

	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		fmt.Println(err)
	}

	return employee, err
}

func GetEmployee(surname, name string) (employee entities.Employee, err error) {
	employee = entities.Employee{}
	row := providers.GetRowEmployee(surname, name)

	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		fmt.Println(err)
	}

	return employee, err
}

func NewEmployee(Surname, Name, Position string) *entities.Employee {
	return &entities.Employee{Surname: Surname, Name: Name, Position: Position}
}

func EmployeeAdd(employee *entities.Employee) (id int, err error) {
	id, err = providers.EmployeeRowAdd(employee)

	return id, err
}

func GetIdDesignatedEmployee(DesignatedEmployee string) (idDesignatedEmployee string) {
	massSurName := strings.Split(DesignatedEmployee, " ")

	surname := massSurName[0]
	name := massSurName[1]

	rowDesignatedEmployee, err := GetEmployee(surname, name)
	if err != nil {
		fmt.Println(err)
	}

	idDesignatedEmployee = strconv.Itoa(rowDesignatedEmployee.Id)

	return idDesignatedEmployee
}
