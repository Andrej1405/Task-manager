package entities

import (
	"app/app/server"
)

type Task struct {
	Id_project         int
	IdTask             int
	Task               string
	DesignatedEmployee int
	Hours              int
	HoursSpent         int
	StatusTask         string
}

func GetAllTasks() (tasks []Task, err error) {
	query := `SELECT * FROM tasksproject`
	rows, err := server.Db.Query(query)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	task := Task{}
	for rows.Next() {
		err = rows.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

func GetTaskById(taskId string) (task Task, err error) {
	row := server.Db.QueryRow(`SELECT * FROM tasksproject WHERE id = $1`, taskId)
	task = Task{}
	err = row.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask)

	if err != nil {
		panic(err)
	}

	return task, err
}

func NewTask(id_project int, task string, designatedEmployee int, hours int, hoursSpent int, statusTask string) *Task {
	return &Task{Id_project: id_project, Task: task, DesignatedEmployee: designatedEmployee, Hours: hours, HoursSpent: hoursSpent, StatusTask: statusTask}
}

func TaskAdd(task *Task) (id int, err error) {
	server.Db.QueryRow(`INSERT INTO tasksproject (id_project, task, designatedemployee, hours, hoursspent, statustask) VALUES ($1, $2, $3, $4, $5, $6) 
	returning id`, task.Id_project, task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask).Scan(&id)

	if err != nil {
		panic(err)
	}

	return id, err
}

func TaskUpdate(task *Task) (err error) {
	_, err = server.Db.Exec(`UPDATE tasksproject SET task = $1, designatedemployee = $2, hours = $3,  hoursspent = $4, statustask = $5 WHERE id = $6`,
		task.Task, task.DesignatedEmployee, task.Hours, task.HoursSpent, task.StatusTask, task.IdTask)
	return
}
