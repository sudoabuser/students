package routes

import (
	"backend/internal/student/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, studentController *controllers.StudentController) {
	router.GET("/students", studentController.GetAll)
	router.GET("/students/:id", studentController.Get)
	router.DELETE("/students/:id", studentController.Delete)
	router.POST("/students", studentController.Add)
}
