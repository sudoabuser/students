package repository

import (
	"backend/internal/student/models"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	GetDB() (*gorm.DB, error)
}
type studentRepository struct {
	DB *gorm.DB
}

func GetDB() (*gorm.DB, error) {
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/fakedatabase?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func NewStudentRepository(db *gorm.DB) (*studentRepository, error) {
	r := &studentRepository{
		DB: db,
	}
	return r, nil
}

func (r *studentRepository) GetAll(page int, pageSize int) ([]models.Student, error) {
	var studentEntities []models.StudentEntity
	offset := (page - 1) * pageSize
	err := r.DB.Offset(offset).Limit(pageSize).Find(&studentEntities).Error

	if err != nil {
		return nil, err
	}

	var students []models.Student
	for _, entity := range studentEntities {
		student := EntityToModel(&entity)
		students = append(students, *student)
	}

	return students, nil
}

func (r *studentRepository) Delete(id uuid.UUID) error {
	entity := &models.StudentEntity{ID: id}
	err := r.DB.Delete(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) Get(id uuid.UUID) (*models.Student, error) {
	var entity models.StudentEntity
	err := r.DB.Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	student := EntityToModel(&entity)
	return student, nil
}

func (r *studentRepository) Add(student *models.Student) error {
	entity := ModelToEntity(student)
	err := r.DB.Create(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func ModelToEntity(student *models.Student) *models.StudentEntity {
	return &models.StudentEntity{
		ID:      uuid.MustParse(student.ID),
		Name:    student.Name,
		Surname: student.Surname,
	}
}

func EntityToModel(entity *models.StudentEntity) *models.Student {
	return &models.Student{
		ID:      entity.ID.String(),
		Name:    entity.Name,
		Surname: entity.Surname,
	}
}

func (r *studentRepository) TotalStudentCount() (int64, error) {
	var totalStudents int64

	err := r.DB.Model(&models.Student{}).Count(&totalStudents).Error
	if err != nil {
		return 0, err
	}
	return totalStudents, nil
}
