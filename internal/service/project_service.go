package service

import (
	"student_pms/internal/models"
	"student_pms/internal/repository"
)

// ProjectService handles business logic related to projects
type ProjectService struct {
	Repo *repository.ProjectRepository
}

// GetProjects retrieves all projects
func (s *ProjectService) GetProjects() ([]models.Project, error) {
	return s.Repo.GetProjects()
}

// AddProject adds a new project
func (s *ProjectService) AddProject(project models.Project) error {
	return s.Repo.AddProject(project)
}

// UpdateProject updates an existing project
func (s *ProjectService) UpdateProject(project models.Project) error {
	return s.Repo.UpdateProject(project)
}

// DeleteProject deletes a project
func (s *ProjectService) DeleteProject(projectID int) error {
	return s.Repo.DeleteProject(projectID)
}
