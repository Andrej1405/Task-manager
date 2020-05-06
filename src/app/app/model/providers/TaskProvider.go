package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
)

type TaskProvider struct {
	mapper   *mappers.TaskMapper
	provider *EmployeeProvider
}

//
func (p *TaskProvider) GetAllTasks() (tasks []entities.Task, err error) {
	p.mapper = new(mappers.TaskMapper)

	tasks, err = p.mapper.GetTasksByDB()
	if err != nil {
		return tasks, err
	}

	return tasks, err
}

func (p *TaskProvider) UpdatingTask(DesignatedEmployee string, Hours int, HoursSpent int, IdTask, Id_project, StatusTask, Task, TaskDescription string) (err error) {
	p.mapper = new(mappers.TaskMapper)
	task, err := p.mapper.GetTaskById(IdTask)

	p.provider = new(EmployeeProvider)
	idDesignatedEmployee, err := p.provider.GetIdDesignatedEmployee(DesignatedEmployee)
	if err != nil {
		fmt.Println(err)
		return err
	}

	task.Task = Task
	task.DesignatedEmployee = idDesignatedEmployee
	task.Hours = Hours
	task.HoursSpent = HoursSpent
	task.StatusTask = StatusTask
	task.TaskDescription = TaskDescription

	err = p.mapper.TaskUpdate(&task)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func (p *TaskProvider) NewTask(id_project int, task, designatedEmployee string, hours int, hoursSpent int, statusTask, taskDescription string) (id int, err error) {
	newTask := &entities.Task{Id_project: id_project, Task: task, DesignatedEmployee: designatedEmployee, Hours: hours,
		HoursSpent: hoursSpent, StatusTask: statusTask, TaskDescription: taskDescription}

	p.mapper = new(mappers.TaskMapper)

	id, err = p.mapper.TaskAdd(newTask)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, err
}
