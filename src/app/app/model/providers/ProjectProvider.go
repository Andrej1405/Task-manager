package providers

import (
	"app/app/model/entities"
	mappers "app/app/model/mappers"
	"fmt"
)

type ProjectProvider struct {
	mapper *mappers.ProjectMapper
}

func (p *ProjectProvider) GetAllProjects() (projects []entities.Project, err error) {
	p.mapper = new(mappers.ProjectMapper)

	projects, err = p.mapper.GetProjectFromBase()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return projects, err
}

func (p *ProjectProvider) NewProject(name string) (id int, err error) {
	p.mapper = new(mappers.ProjectMapper)

	project := &entities.Project{Name: name}

	id, err = p.mapper.ProjectAdd(project)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, err
}

func (p *ProjectProvider) UpdateProj(id, name string) (err error) {
	p.mapper = new(mappers.ProjectMapper)

	project, err := p.mapper.GetProjectById(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	project.Name = name

	err = p.mapper.ProjectUpdate(&project)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func (p *ProjectProvider) DeleteProjectById(id string) (err error) {
	p.mapper = new(mappers.ProjectMapper)

	err = p.mapper.ProjectDelete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}
