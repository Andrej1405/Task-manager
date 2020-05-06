package controllers

import (
	"app/app/model/providers"

	"github.com/revel/revel"
)

type ControllerTask struct {
	*revel.Controller
	providerTask     *providers.TaskProvider
	providerEmployee *providers.EmployeeProvider
}

// Получение задач и возвращение их на фронт.
func (c ControllerTask) GetTask() revel.Result {
	c.providerTask = new(providers.TaskProvider)

	tasks, err := c.providerTask.GetAllTasks()
	if err != nil {
		return c.RenderError(err)
	}

	return c.RenderJSON(tasks)
}

// Добавление новой задачи. Возвращает id новой задачи на фронт.
func (c ControllerTask) AddNewTask(Task, DesignatedEmployee string, Hours, HoursSpent int, StatusTask, TaskDescription string, Id_project int) revel.Result {
	c.providerEmployee = new(providers.EmployeeProvider)

	idDesignatedEmployee, err := c.providerEmployee.GetIdDesignatedEmployee(DesignatedEmployee)
	if err != nil {
		return c.RenderError(err)
	}

	c.providerTask = new(providers.TaskProvider)
	id, err := c.providerTask.NewTask(Id_project, Task, idDesignatedEmployee, Hours, HoursSpent, StatusTask, TaskDescription)
	if err != nil {
		return c.RenderError(err)
	}

	return c.RenderJSON(id)
}

// Обновление существующей задачи.
func (c ControllerTask) UpdateTask(DesignatedEmployee string, Hours, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) revel.Result {
	c.providerTask = new(providers.TaskProvider)

	err := c.providerTask.UpdatingTask(DesignatedEmployee, Hours, HoursSpent, IdTask, Id_project, StatusTask, Task, TaskDescription)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Render()
}
