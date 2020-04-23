package entities

import (
	"app/app/server"
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
		return employees, err
	}
	defer rows.Close()

	employee := Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
		if err != nil {
			return employees, err
		}
		employees = append(employees, employee)
	}

	return employees, err
}

func GetEmployeeById(employeeId string) (employee Employee, err error) {
	row := server.Db.QueryRow(`SELECT * FROM employees WHERE id = $1`, employeeId)
	employee = Employee{}
	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		panic(err)
	}
	return employee, err
}

func NewEmployee(Surname, Name, Position string) *Employee {
	return &Employee{Surname: Surname, Name: Name, Position: Position}
}

func EmployeeAdd(employee *Employee) (id int, err error) {
	server.Db.QueryRow(`INSERT INTO employees (surname, name, position) VALUES ($1, $2, $3) returning id`, employee.Surname, employee.Name, employee.Position).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, err
}

func EmployeeUpdate(employee *Employee) (err error) {
	_, err = server.Db.Exec(`UPDATE employees SET surname = $1, name = $2, position = $3 WHERE id = $4`, employee.Surname, employee.Name,
		employee.Position, employee.Id)
	return
}

func EmployeeDelete(employeeId string) (err error) {
	_, err = server.Db.Exec(`DELETE FROM employees WHERE id = $1`, employeeId)
	if err != nil {
		panic(err)
	}
	return
}
