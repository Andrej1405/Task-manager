package controllers

import (
	entities "app/app/model"
	"app/app/server"

	"github.com/revel/revel"
)

type ControllerTask struct {
	*revel.Controller
}

func (c ControllerTask) GetTask() revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	tasks, err := entities.GetAllTasks()
	if err != nil {
		panic(err)
	}

	return c.RenderJSON(tasks)
}

func (c ControllerTask) AddNewTask(Task string, DesignatedEmployee int, Hours int, HoursSpent int, StatusTask string, Id_project int) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	task := entities.NewTask(Id_project, Task, DesignatedEmployee, Hours, HoursSpent, StatusTask)

	id, err := entities.TaskAdd(task)
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(id)
}

func (c ControllerTask) DeleteTask() revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	return c.Render()
}

func (c ControllerTask) UpdateTask(DesignatedEmployee int, Hours int, HoursSpent int, IdTask, Id_project, StatusTask, Task string) revel.Result {
	err := server.InitDB()
	if err != nil {
		panic(err)
	}

	task, err := entities.GetTaskById(IdTask)
	if err != nil {
		panic(err)
	}

	task.Task = Task
	task.DesignatedEmployee = DesignatedEmployee
	task.Hours = Hours
	task.HoursSpent = HoursSpent
	task.StatusTask = StatusTask

	err = entities.TaskUpdate(&task)
	if err != nil {
		panic(err)
	}

	return c.Render()
}
