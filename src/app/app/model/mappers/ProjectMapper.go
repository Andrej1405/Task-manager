package mappers

import (
	"app/app/model/entities"
	"app/app/model/providers"
	"fmt"
)

func GetAllProjects() (projects []entities.Project, err error) {
	project := entities.Project{}
	rows := providers.GetAllRowsProject()

	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, project)
	}

	return projects, err
}

func GetProjectById(projectId string) (project entities.Project, err error) {
	project = entities.Project{}
	row := providers.GetRowProjectById(projectId)

	err = row.Scan(&project.Id, &project.Name)
	if err != nil {
		fmt.Println(err)
	}

	return project, err
}

func NewProject(name string) *entities.Project {
	return &entities.Project{Name: name}
}

func ProjectAdd(project *entities.Project) (id int, err error) {
	id, err = providers.AddRowProject(project)

	return id, err
}
