package entities

import (
	"app/app/server"
	"fmt"
)

type Project struct {
	Id   int
	Name string
}

func GetAllProjects() (projects []Project, err error) {
	query := `SELECT * FROM projects`
	rows, err := server.Db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	project := Project{}
	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, project)
	}

	return projects, err
}

func GetProjectById(projectId string) (project Project, err error) {
	query := `SELECT * FROM projects WHERE id = $1`
	row := server.Db.QueryRow(query, projectId)

	project = Project{}

	err = row.Scan(&project.Id, &project.Name)
	if err != nil {
		fmt.Println(err)
	}

	return project, err
}

func NewProject(name string) *Project {
	return &Project{Name: name}
}

func ProjectAdd(project *Project) (id int, err error) {
	query := `INSERT INTO projects (name) VALUES ($1) returning id`

	server.Db.QueryRow(query, project.Name).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func ProjectUpdate(project *Project) (err error) {
	query := `UPDATE projects SET name = $1 WHERE id = $2`

	_, err = server.Db.Exec(query, project.Name, project.Id)

	return
}

func ProjectDelete(projectId string) (err error) {
	query := `DELETE FROM projects WHERE id = $1`
	_, err = server.Db.Exec(query, projectId)

	if err != nil {
		fmt.Println(err)
	}

	return
}
