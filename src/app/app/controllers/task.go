package controllers

import (
	entities "app/app/model"
	"app/app/server"
	"fmt"

	"github.com/revel/revel"
)

type ControllerTask struct {
	*revel.Controller
}

func (c ControllerTask) GetTask() revel.Result {
	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	tasks, err := entities.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(tasks)
}

func (c ControllerTask) AddNewTask(Task string, DesignatedEmployee, Hours, HoursSpent int, StatusTask string, TaskDescription string, Id_project int) revel.Result {
	if Task == "" || DesignatedEmployee == 0 || Hours == 0 || HoursSpent == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == 0 {
		fmt.Println("Данные введены некорректно. Добавить сотрудника не удалось")
		c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	task := entities.NewTask(Id_project, Task, DesignatedEmployee, Hours, HoursSpent, StatusTask, TaskDescription)

	id, err := entities.TaskAdd(task)
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(id)
}

// func (c ControllerTask) DeleteTask() revel.Result {
// 	err := server.InitDB()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return c.Render()
// }

func (c ControllerTask) UpdateTask(DesignatedEmployee int, Hours int, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) revel.Result {
	if Task == "" || DesignatedEmployee == 0 || Hours == 0 || HoursSpent == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == "" || IdTask == "" {
		fmt.Println("Данные введены некорректно. Обновить информацию по сотруднику не удалось")
		c.Render()
	}

	err := server.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	task, err := entities.GetTaskById(IdTask)
	if err != nil {
		fmt.Println(err)
	}

	task.Task = Task
	task.DesignatedEmployee = DesignatedEmployee
	task.Hours = Hours
	task.HoursSpent = HoursSpent
	task.StatusTask = StatusTask
	task.TaskDescription = TaskDescription

	err = entities.TaskUpdate(&task)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
