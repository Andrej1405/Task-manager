package entities

import (
	"app/app/server"
	"fmt"
)

type Task struct {
	Id_project         int
	IdTask             int
	Task               string
	DesignatedEmployee int
	Hours              int
	HoursSpent         int
	StatusTask         string
	TaskDescription    string
}

func GetAllTasks() (tasks []Task, err error) {
	query := `SELECT * FROM tasksproject`
	rows, err := server.Db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	task := Task{}
	for rows.Next() {
		err = rows.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask, &task.TaskDescription)
		if err != nil {
			fmt.Println(err)
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

func GetTaskById(taskId string) (task Task, err error) {
	query := `SELECT * FROM tasksproject WHERE id = $1`
	row := server.Db.QueryRow(query, taskId)

	task = Task{}

	err = row.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask, &task.TaskDescription)

	if err != nil {
		fmt.Println(err)
	}

	return task, err
}

func NewTask(id_project int, task string, designatedEmployee int, hours int, hoursSpent int, statusTask string, taskDescription string) *Task {
	return &Task{Id_project: id_project, Task: task, DesignatedEmployee: designatedEmployee, Hours: hours,
		HoursSpent: hoursSpent, StatusTask: statusTask, TaskDescription: taskDescription}
}

func TaskAdd(task *Task) (id int, err error) {
	query := `INSERT INTO tasksproject (id_project, task, designatedemployee, hours, hoursspent, statustask, taskdescription) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	returning id`
	server.Db.QueryRow(query, task.Id_project, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.TaskDescription).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func TaskUpdate(task *Task) (err error) {
	query := `UPDATE tasksproject SET task = $1, designatedemployee = $2, hours = $3,  hoursspent = $4, statustask = $5, taskdescription = $6 WHERE id = $7`
	_, err = server.Db.Exec(query, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.TaskDescription, task.IdTask)

	return
}
