package main

import (
	"backend/internal/student/controllers"
	"backend/internal/student/models"
	"backend/internal/student/repository"
	"backend/internal/student/routes"
	"backend/internal/student/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.GetDB()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	repo, err := repository.NewStudentRepository(db)
	if err != nil {
		log.Fatal("variable 'repo' couldn't be initialized", err)
	}
	Service := services.Service(repo)
	Controller := controllers.Controller(Service)

	db.AutoMigrate(&models.Student{}, &models.StudentEntity{})

	router := gin.Default()
	router.Use(cors.Default())
	routes.SetupRoutes(router, Controller)
	router.Run("localhost:8080")
}
