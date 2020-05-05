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
	tasks, err := mappers.GetAllTask()
	if err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(tasks)
}

func (c ControllerTask) AddNewTask(Task, DesignatedEmployee string, Hours, HoursSpent int, StatusTask, TaskDescription string, Id_project int) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == 0 {
		fmt.Println("Данные введены некорректно. Добавить сотрудника не удалось")
		return c.Render()
	}

	idDesignatedEmployee := providers.GetIdDesignatedEmployee(DesignatedEmployee)

	id := providers.NewTask(Id_project, Task, idDesignatedEmployee, Hours, HoursSpent, StatusTask, TaskDescription)

	return c.RenderJSON(id)
}

func (c ControllerTask) UpdateTask(DesignatedEmployee string, Hours, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) revel.Result {
	if Task == "" || DesignatedEmployee == "" || Hours == 0 || StatusTask == "" || TaskDescription == "" ||
		Id_project == "" || IdTask == "" {
		fmt.Println("Данные введены некорректно. Обновить информацию по сотруднику не удалось")
		return c.Render()
	}

	err := providers.UpdatingTask(DesignatedEmployee, Hours, HoursSpent, IdTask, Id_project, StatusTask, Task, TaskDescription)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render()
}
