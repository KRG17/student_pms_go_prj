// docker run --name student-project-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=student_project_db -p 3306:3306 -d mysql:8.0.30

package repository

import (
	"database/sql"
	"student_pms/internal/models"
)

// ProjectRepository handles database operations related to projects
type ProjectRepository struct {
	DB *sql.DB
}

// GetProjects retrieves all projects from the database
func (r *ProjectRepository) GetProjects() ([]models.Project, error) {
	// Implement your database query logic here
	rows, err := r.DB.Query("SELECT id, name, description FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

// AddProject adds a new project to the database
func (r *ProjectRepository) AddProject(project models.Project) error {
	// Implement your database insert logic here
	_, err := r.DB.Exec("INSERT INTO projects (id, name, description) VALUES (?, ?, ?)", project.ID, project.Name, project.Description)
	return err
}

// UpdateProject updates an existing project in the database
func (r *ProjectRepository) UpdateProject(project models.Project) error {
	// Implement your database update logic here
	_, err := r.DB.Exec("UPDATE projects SET name=?, description=? WHERE id=?", project.Name, project.Description, project.ID)
	return err
}

// DeleteProject deletes a project from the database
func (r *ProjectRepository) DeleteProject(projectID int) error {
	// Implement your database delete logic here
	_, err := r.DB.Exec("DELETE FROM projects WHERE id=?", projectID)
	return err
}
