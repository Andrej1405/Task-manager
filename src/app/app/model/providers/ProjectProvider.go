package providers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

func GetAllRowsProject() *sql.Rows {
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

	return rows
}

func GetRowProjectById(projectId string) *sql.Row {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM projects WHERE id = $1`
	row := db.QueryRow(query, projectId)

	return row
}

func AddRowProject(project *entities.Project) (id int, err error) {
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

func ProjectRowUpdate(project *entities.Project) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `UPDATE projects SET name = $1 WHERE id = $2`

	_, err = db.Exec(query, project.Name, project.Id)

	return
}

func ProjectRowDelete(projectId string) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `DELETE FROM projects WHERE id = $1`
	_, err = db.Exec(query, projectId)

	if err != nil {
		fmt.Println(err)
	}

	return
}
