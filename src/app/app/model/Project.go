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

func GetProjectById(projectId string) (project Project, err error) {
	row := server.Db.QueryRow(`SELECT * FROM projects WHERE id = $1`, projectId)
	project = Project{}
	err = row.Scan(&project.Id, &project.Name, &project.Date)
	if err != nil {
		panic(err)
	}
	return project, err
}

func NewProject(name, date string) *Project {
	return &Project{Name: name, Date: date}
}

func ProjectAdd(project *Project) (id int, err error) {
	server.Db.QueryRow(`INSERT INTO projects (name, date) VALUES ($1, $2) returning id`, project.Name, project.Date).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, err
}

// func ProjectUpdate(project *Project) (err error) {
// 	_, err = server.Db.Exec(`UPDATE project SET name = $1, date = $2 WHERE id = $3`, project.Name, project.Date, project.Id)
// 	return
// }

func ProjectDelete(projectId string) (err error) {
	_, err = server.Db.Exec(`DELETE FROM projects WHERE id = $1`, projectId)
	if err != nil {
		panic(err)
	}
	return
}
