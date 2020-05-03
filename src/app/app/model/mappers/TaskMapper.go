package mappers

import (
	"app/app/model/entities"
	"app/app/model/providers"
	"fmt"
)

func GetAllTasks() (tasks []entities.Task, err error) {
	task := entities.Task{}
	rows := providers.GetAllRowsTask()

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

func GetTaskById(taskId string) (task entities.Task, err error) {
	task = entities.Task{}
	row := providers.GetRowTaskById(taskId)

	err = row.Scan(&task.Id_project, &task.IdTask, &task.Task, &task.DesignatedEmployee, &task.Hours, &task.HoursSpent, &task.StatusTask, &task.TaskDescription)

	if err != nil {
		fmt.Println(err)
	}

	return task, err
}

func NewTask(id_project int, task string, designatedEmployee string, hours int, hoursSpent int, statusTask string, taskDescription string) *entities.Task {
	return &entities.Task{Id_project: id_project, Task: task, DesignatedEmployee: designatedEmployee, Hours: hours,
		HoursSpent: hoursSpent, StatusTask: statusTask, TaskDescription: taskDescription}
}

func TaskAdd(task *entities.Task) (id int, err error) {
	id, err = providers.TaskRowAdd(task)

	return id, err
}
