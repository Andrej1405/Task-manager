package mappers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

func GetAllTask() (tasks []entities.Task, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM tasksproject`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	task := entities.Task{}

	for rows.Next() {
		err = rows.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask, &task.TaskDescription)
		if err != nil {
			fmt.Println(err)
		}
		designatedEmployee, err := GetEmployeeById(task.DesignatedEmployee)
		if err != nil {
			fmt.Println(err)
		}
		task.DesignatedEmployee = designatedEmployee.Surname + " " + designatedEmployee.Name
		tasks = append(tasks, task)
	}

	return tasks, err
}

func GetTaskById(id string) (task entities.Task, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM tasksproject WHERE id = $1`
	row := db.QueryRow(query, id)

	task = entities.Task{}

	err = row.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask, &task.TaskDescription)

	if err != nil {
		fmt.Println(err)
	}

	return task, err
}

func TaskUpdate(task *entities.Task) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `UPDATE tasksproject SET task = $1, designatedemployee = $2, hours = $3,  hoursspent = $4, statustask = $5, taskdescription = $6 WHERE id = $7`
	_, err = db.Exec(query, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.TaskDescription, task.IdTask)

	return
}

func TaskAdd(task *entities.Task) (id int, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `INSERT INTO tasksproject (id_project, task, designatedemployee, hours, hoursspent, statustask, taskdescription) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	returning id`
	db.QueryRow(query, task.Id_project, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.TaskDescription).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	return id, err
}
