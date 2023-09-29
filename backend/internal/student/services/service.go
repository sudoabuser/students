package services

import (
	"backend/internal/student/models"
	"errors"

	"github.com/google/uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*models.Student, error)
	Add(student *models.Student) error
	Delete(id uuid.UUID) error
	GetAll(page int, pageSize int) ([]models.Student, error)
	TotalStudentCount() (int64, error)
}

type StudentService struct {
	repository Repository
}

func Service(repository Repository) *StudentService {
	return &StudentService{repository: repository}
}

func (s *StudentService) Get(id uuid.UUID) (*models.Student, error) {
	student, err := s.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) Delete(id uuid.UUID) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func (s *StudentService) Add(student *models.Student) error {
	if student.Name == "" || student.Surname == "" {
		return errors.New("name and surname are required") //bunlari sanirim controller'a almaliyim??
	}
	uID := uuid.New()
	student.ID = uID.String()

	err := s.repository.Add(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentService) GetAll(page int, pageSize int) (models.PaginationResponse, error) {
	if page <= 0 || pageSize <= 0 {
		return models.PaginationResponse{}, errors.New("page and pagesize cannot be lower than 1") //bunlari sanirim controller'a almaliyim??
	}
	students, err := s.repository.GetAll(page, pageSize)
	if err != nil {
		return models.PaginationResponse{}, err
	}

	totalStudents, err := s.repository.TotalStudentCount()
	if err != nil {
		return models.PaginationResponse{}, err
	}

	totalPages := (totalStudents + int64(pageSize) - 1) / int64(pageSize)

	pageInfo := models.Page{
		Number:   page,
		Size:     pageSize,
		Elements: int(totalStudents),
		Pages:    int(totalPages),
	}

	response := models.PaginationResponse{
		Students: students,
		Page:     pageInfo,
	}

	return response, nil
}
