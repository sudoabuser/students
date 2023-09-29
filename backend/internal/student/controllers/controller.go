package controllers

import (
	"backend/internal/student/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentService interface {
	GetAll(page int, pageSize int) (models.PaginationResponse, error)
	Get(id uuid.UUID) (*models.Student, error)
	Add(student *models.Student) error
	Delete(id uuid.UUID) error
}

type StudentController struct {
	Service StudentService
}

func Controller(Service StudentService) *StudentController {
	return &StudentController{Service: Service}
}

func (c *StudentController) Get(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "invalid UUID"})
		return
	}

	student, err := c.Service.Get(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "student not found", "student_id": id})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	err = c.Service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func (c *StudentController) Add(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	err := c.Service.Add(&student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}
	ctx.JSON(http.StatusCreated, student)
}

func (c *StudentController) GetAll(ctx *gin.Context) {
	page := 1
	pageSize := 10
	if pageStr := ctx.Query("page"); pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if pageSizeStr := ctx.Query("size"); pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	response, err := c.Service.GetAll(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retreive students"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
