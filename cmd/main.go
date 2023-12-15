package main

import (
	"database/sql"
	"log"
	"student_pms/internal/handlers"
	"student_pms/internal/repository"
	"student_pms/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize MySQL database connection
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/student_project_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	projectRepo := &repository.ProjectRepository{DB: db}
	projectService := &service.ProjectService{Repo: projectRepo}
	projectHandler := &handlers.ProjectHandler{Service: projectService}

	// Set up the HTTP server
	router := gin.Default()

	// Define routes
	router.GET("/projects", projectHandler.GetProjectsHandler)
	router.POST("/projects", projectHandler.AddProjectHandler)
	router.PUT("/projects/:id", projectHandler.UpdateProjectHandler)
	router.DELETE("/projects/:id", projectHandler.DeleteProjectHandler)

	// Start the server on port 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
