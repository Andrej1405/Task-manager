package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
)

func UpdateProj(id, name string) (err error) {
	project := mappers.GetProjectById(id)
	project.Name = name

	err = mappers.ProjectUpdate(&project)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func NewProject(name string) (id int) {
	project := &entities.Project{Name: name}

	id, err := mappers.ProjectAdd(project)
	if err != nil {
		fmt.Println(err)
	}

	return id
}
