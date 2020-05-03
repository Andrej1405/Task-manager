package controllers

import (
	mappers "app/app/model/mappers"
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerTask struct {
	*revel.Controller
}

func (c ControllerTask) GetTask() revel.Result {
	tasks, err := mappers.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(tasks)
}

func (c ControllerTask) AddNewTask(Task string, DesignatedEmployee string, Hours, HoursSpent int, StatusTask string, TaskDescription string, Id_project int) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || HoursSpent == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == 0 {
		fmt.Println("Данные введены некорректно. Добавить сотрудника не удалось")
		c.Render()
	}

	idDesignatedEmployee := mappers.GetIdDesignatedEmployee(DesignatedEmployee)

	task := mappers.NewTask(Id_project, Task, idDesignatedEmployee, Hours, HoursSpent, StatusTask, TaskDescription)

	id, err := mappers.TaskAdd(task)
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(id)
}

func (c ControllerTask) UpdateTask(DesignatedEmployee string, Hours int, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || HoursSpent == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == "" || IdTask == "" {
		fmt.Println("Данные введены некорректно. Обновить информацию по сотруднику не удалось")
		c.Render()
	}

	task, err := mappers.GetTaskById(IdTask)
	if err != nil {
		fmt.Println(err)
	}

	idDesignatedEmployee := mappers.GetIdDesignatedEmployee(DesignatedEmployee)

	task.Task = Task
	task.DesignatedEmployee = idDesignatedEmployee
	task.Hours = Hours
	task.HoursSpent = HoursSpent
	task.StatusTask = StatusTask
	task.TaskDescription = TaskDescription

	err = providers.TaskRowUpdate(&task)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
