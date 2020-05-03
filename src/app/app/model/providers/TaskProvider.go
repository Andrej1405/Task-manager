package providers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

func GetAllRowsTask() *sql.Rows {
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

	return rows
}

func GetRowTaskById(taskId string) *sql.Row {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM tasksproject WHERE id = $1`
	row := db.QueryRow(query, taskId)

	return row
}

func TaskRowUpdate(task *entities.Task) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `UPDATE tasksproject SET task = $1, designatedemployee = $2, hours = $3,  hoursspent = $4, statustask = $5, taskdescription = $6 WHERE id = $7`
	_, err = db.Exec(query, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.TaskDescription, task.IdTask)

	return
}

func TaskRowAdd(task *entities.Task) (id int, err error) {
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
