package controllers

import (
	"app/app/model/providers"
	"fmt"

	"github.com/revel/revel"
)

type ControllerTask struct {
	*revel.Controller
	provider         *providers.TaskProvider
	providerEmployee *providers.EmployeeProvider
}

// Получение задач и возвращение их на фронт.
func (c ControllerTask) GetTask() revel.Result {
	c.provider = new(providers.TaskProvider)

	tasks, err := c.provider.GetAllTasks()
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.RenderJSON(tasks)
}

// Добавление новой задачи. Возвращает id новой задачи на фронт.
func (c ControllerTask) AddNewTask(Task, DesignatedEmployee string, Hours, HoursSpent int, StatusTask, TaskDescription string, Id_project int) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == 0 {
		fmt.Println("Данные введены некорректно. Добавить сотрудника не удалось")
		return c.Render()
	}

	c.providerEmployee = new(providers.EmployeeProvider)
	idDesignatedEmployee, err := c.providerEmployee.GetIdDesignatedEmployee(DesignatedEmployee)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	c.provider = new(providers.TaskProvider)
	id, err := c.provider.NewTask(Id_project, Task, idDesignatedEmployee, Hours, HoursSpent, StatusTask, TaskDescription)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.RenderJSON(id)
}

// Обновление существующей задачи.
func (c ControllerTask) UpdateTask(DesignatedEmployee string, Hours, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == "" || IdTask == "" {
		fmt.Println("Данные введены некорректно. Обновить информацию по сотруднику не удалось")
		return c.Render()
	}

	c.provider = new(providers.TaskProvider)

	err := c.provider.UpdatingTask(DesignatedEmployee, Hours, HoursSpent, IdTask, Id_project, StatusTask, Task, TaskDescription)
	if err != nil {
		fmt.Println(err)
		return c.RenderError(err)
	}

	return c.Render()
}
