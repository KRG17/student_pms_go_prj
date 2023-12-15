package handlers

import (
	"net/http"
	"strconv"
	"student_pms/internal/models"
	"student_pms/internal/service"

	"github.com/gin-gonic/gin"
)

// ProjectHandler handles HTTP requests related to projects
type ProjectHandler struct {
	Service *service.ProjectService
}

// GetProjectsHandler handles the GET request for projects
func (h *ProjectHandler) GetProjectsHandler(c *gin.Context) {
	projects, err := h.Service.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// AddProjectHandler handles the POST request to add a project
func (h *ProjectHandler) AddProjectHandler(c *gin.Context) {
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.Service.AddProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// UpdateProjectHandler handles the PUT request to update a project
func (h *ProjectHandler) UpdateProjectHandler(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var updatedProject models.Project
	if err := c.BindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedProject.ID = projectID

	if err := h.Service.UpdateProject(updatedProject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, updatedProject)
}

// DeleteProjectHandler handles the DELETE request to delete a project
func (h *ProjectHandler) DeleteProjectHandler(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Delete the project
	if err := h.Service.DeleteProject(projectID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Fetch the remaining projects
	projects, err := h.Service.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, projects)
}
