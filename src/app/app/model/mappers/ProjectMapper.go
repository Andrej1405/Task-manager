package mappers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

type ProjectMapper struct {
	Project *entities.Project
}

// Получение всех проектов из базы данных.
func (m *ProjectMapper) GetProjectFromBase() (projects []entities.Project, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	query := `SELECT * FROM projects`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	project := entities.Project{}

	for rows.Next() {
		err = rows.Scan(&project.Id, &project.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, err
}

// Получение проекта по его id из базы данных.
func (m *ProjectMapper) GetProjectById(id string) (project entities.Project, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
		return project, err
	}
	defer db.Close()

	query := `SELECT * FROM projects WHERE id = $1`
	row := db.QueryRow(query, id)

	project = entities.Project{}

	err = row.Scan(&project.Id, &project.Name)
	if err != nil {
		fmt.Println(err)
		return project, err
	}

	return project, err
}

// Добавление нового проекта.
func (m *ProjectMapper) ProjectAdd(project *entities.Project) (id int, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer db.Close()

	query := `INSERT INTO projects (name) VALUES ($1) returning id`

	err = db.QueryRow(query, project.Name).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, err
}

// Обновление информации по проекту.
func (m *ProjectMapper) ProjectUpdate(project *entities.Project) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	query := `UPDATE projects SET name = $1 WHERE id = $2`

	_, err = db.Exec(query, project.Name, project.Id)

	return
}

// Удаление проекта.
func (m *ProjectMapper) ProjectDelete(id string) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	query := `DELETE FROM projects WHERE id = $1`
	_, err = db.Exec(query, id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}
