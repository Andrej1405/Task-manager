package mappers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

func GetAllProject() (projects []entities.Project, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM projects`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	project := entities.Project{}

	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			fmt.Println(err)
		}
		projects = append(projects, project)
	}

	return projects, err
}

func GetProjectById(id string) (project entities.Project) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM projects WHERE id = $1`
	row := db.QueryRow(query, id)

	project = entities.Project{}

	err = row.Scan(&project.Id, &project.Name)
	if err != nil {
		fmt.Println(err)
	}

	return project
}

func ProjectAdd(project *entities.Project) (id int, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `INSERT INTO projects (name) VALUES ($1) returning id`

	db.QueryRow(query, project.Name).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func ProjectUpdate(project *entities.Project) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `UPDATE projects SET name = $1 WHERE id = $2`

	_, err = db.Exec(query, project.Name, project.Id)

	return
}

func ProjectDelete(id string) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `DELETE FROM projects WHERE id = $1`
	_, err = db.Exec(query, id)

	if err != nil {
		fmt.Println(err)
	}

	return
}
