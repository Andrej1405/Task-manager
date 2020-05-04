package mappers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

func GetAllEmployees() (employees []entities.Employee, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM employees`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	employee := entities.Employee{}

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
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM employees WHERE id = $1`
	row := db.QueryRow(query, employeeId)

	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		fmt.Println(err)
	}

	return employee, err
}

func GetEmployeeBySurnameName(surname, name string) (employee entities.Employee, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM employees WHERE (surname = $1 AND name = $2)`
	row := db.QueryRow(query, surname, name)

	employee = entities.Employee{}

	err = row.Scan(&employee.Id, &employee.Surname, &employee.Name, &employee.Position)
	if err != nil {
		fmt.Println(err)
	}

	return employee, err
}

func EmployeeAdd(employee *entities.Employee) (id int, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `INSERT INTO employees (surname, name, position) VALUES ($1, $2, $3) returning id`
	db.QueryRow(query, employee.Surname, employee.Name, employee.Position).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func EmployeeUpdate(employee *entities.Employee) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `UPDATE employees SET surname = $1, name = $2, position = $3 WHERE id = $4`
	_, err = db.Exec(query, employee.Surname, employee.Name,
		employee.Position, employee.Id)

	return
}

func EmployeeDelete(employeeId string) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `DELETE FROM employees WHERE id = $1`
	_, err = db.Exec(query, employeeId)
	if err != nil {
		fmt.Println(err)
	}

	return
}
