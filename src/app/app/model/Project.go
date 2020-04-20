package entities

import (
	"app/app/server"
)

type Project struct {
	Id   int
	Name string
	Date string
}

func GetAllProjects() (projects []Project, err error) {
	query := `SELECT * FROM projects`
	rows, err := server.Db.Query(query)
	if err != nil {
		return projects, err
	}
	defer rows.Close()

	project := Project{}
	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Name, &project.Date)
		if err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	return projects, err
}
