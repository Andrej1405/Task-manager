package providers

import (
	"app/app/model/entities"
	"app/app/model/mappers"
	"fmt"
)

type ProjectProvider struct {
	mapper *mappers.ProjectMapper
}

func (p *ProjectProvider) UpdateProj(id, name string) (err error) {
	project := p.mapper.GetProjectById(id)
	project.Name = name

	err = p.mapper.ProjectUpdate(&project)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (p *ProjectProvider) NewProject(name string) (id int) {
	project := &entities.Project{Name: name}

	id, err := p.mapper.ProjectAdd(project)
	if err != nil {
		fmt.Println(err)
	}

	return id
}
