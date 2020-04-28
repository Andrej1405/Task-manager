package entities

import (
	"app/app/server"
	"fmt"
)

type Employee struct {
	Id       int
	Surname  string
	Name     string
	Position string
}

func GetAllEmployees() (employees []Employee, err error) {
	query := `SELECT * FROM employees`
	rows, err := server.Db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	employee := Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
		if err != nil {
			fmt.Println(err)
		}
		employees = append(employees, employee)
	}

	return employees, err
}

func GetEmployeeById(employeeId string) (employee Employee, err error) {
	query := `SELECT * FROM employees WHERE id = $1`
	row := server.Db.QueryRow(query, employeeId)

	employee = Employee{}

	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		fmt.Println(err)
	}

	return employee, err
}

func NewEmployee(Surname, Name, Position string) *Employee {
	return &Employee{Surname: Surname, Name: Name, Position: Position}
}

func EmployeeAdd(employee *Employee) (id int, err error) {
	query := `INSERT INTO employees (surname, name, position) VALUES ($1, $2, $3) returning id`
	server.Db.QueryRow(query, employee.Surname, employee.Name, employee.Position).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func EmployeeUpdate(employee *Employee) (err error) {
	query := `UPDATE employees SET surname = $1, name = $2, position = $3 WHERE id = $4`
	_, err = server.Db.Exec(query, employee.Surname, employee.Name,
		employee.Position, employee.Id)
	return
}

func EmployeeDelete(employeeId string) (err error) {
	query := `DELETE FROM employees WHERE id = $1`
	_, err = server.Db.Exec(query, employeeId)
	if err != nil {
		fmt.Println(err)
	}

	return
}
