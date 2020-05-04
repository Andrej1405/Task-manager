package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
)

func UpdatingTask(DesignatedEmployee string, Hours int, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) (err error) {
	task, err := mappers.GetTaskById(IdTask)

	idDesignatedEmployee := GetIdDesignatedEmployee(DesignatedEmployee)

	task.Task = Task
	task.DesignatedEmployee = idDesignatedEmployee
	task.Hours = Hours
	task.HoursSpent = HoursSpent
	task.StatusTask = StatusTask
	task.TaskDescription = TaskDescription

	err = mappers.TaskUpdate(&task)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func NewTask(id_project int, task string, designatedEmployee string, hours int, hoursSpent int, statusTask string, taskDescription string) (id int) {
	newTask := &entities.Task{Id_project: id_project, Task: task, DesignatedEmployee: designatedEmployee, Hours: hours,
		HoursSpent: hoursSpent, StatusTask: statusTask, TaskDescription: taskDescription}

	id, err := mappers.TaskAdd(newTask)
	if err != nil {
		fmt.Println(err)
	}

	return id
}
