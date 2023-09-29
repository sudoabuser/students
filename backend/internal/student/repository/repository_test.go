package repository

import (
	"backend/internal/student/models"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAdd(t *testing.T) {
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/teststudent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo, err := NewStudentRepository(db)
	if err != nil {
		t.Fatalf("Failed to create student repository: %v", err)
	}

	t.Run("AddSuccess", func(t *testing.T) {
		// Get the initial count of students in the database
		initialCount, err := repo.TotalStudentCount()
		fmt.Println("initial student count:", initialCount)
		if err != nil {
			t.Fatalf("Failed to get initial student count: %v", err)
		}
		testStudent := &models.Student{
			ID:      uuid.New().String(),
			Name:    "kamil",
			Surname: "koc",
		}
		err = repo.Add(testStudent)
		assert.NoError(t, err, "Expected Add to succeed, but it didn't")

		//Get the final count of students in the database
		finalCount, err := repo.TotalStudentCount()
		fmt.Println("final student count:", finalCount)

		if err != nil {
			t.Fatalf("Failed to get final student count: %v", err)
		}

		assert.NotEqual(t, initialCount, finalCount, "Student expected to be added, but it didn't")
	})
}

func TestGet(t *testing.T) {
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/teststudent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo, err := NewStudentRepository(db)
	if err != nil {
		t.Fatalf("Failed to create student repository: %v", err)
	}

	expectedStudent := &models.StudentEntity{
		ID:      uuid.New(),
		Name:    "hasan",
		Surname: "huseyin",
	}

	newStudent := EntityToModel(expectedStudent)

	repo.Add(newStudent)

	t.Run("GetExistingStudent", func(t *testing.T) {
		studentEntity := ModelToEntity(newStudent)
		student, err := repo.Get(studentEntity.ID)
		assert.NoError(t, err, "Expected to fetch a student but it didn't")
		assert.NotNil(t, student)
		assert.Equal(t, expectedStudent.ID, studentEntity.ID)
	})

	t.Run("GetNonExistingStudent", func(t *testing.T) {
		student, err := repo.Get(uuid.New())
		assert.Error(t, err)
		assert.Nil(t, student)
	})
}

func TestGetAllStudent(t *testing.T) {
	TestCleanup(t)
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/teststudent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo, err := NewStudentRepository(db)
	if err != nil {
		t.Fatalf("Failed to create student repository: %v", err)
	}

	studentsToAdd := []models.StudentEntity{
		{
			ID:      uuid.New(),
			Name:    "ahmet",
			Surname: "ceylan",
		},
		{
			ID:      uuid.New(),
			Name:    "hasan",
			Surname: "huseyin",
		},
	}

	for _, student := range studentsToAdd {
		newStudent := EntityToModel(&student)
		repo.Add(newStudent)
	}

	testCases := []struct {
		Page        int
		PageSize    int
		Expected    []models.Student
		Description string
	}{
		{
			Page:     1,
			PageSize: 10,
			Expected: []models.Student{
				{
					ID:      studentsToAdd[0].ID.String(),
					Name:    studentsToAdd[0].Name,
					Surname: studentsToAdd[0].Surname,
				},
				{
					ID:      studentsToAdd[1].ID.String(),
					Name:    studentsToAdd[1].Name,
					Surname: studentsToAdd[1].Surname,
				},
			},
			Description: "Get the first page of students",
		},
		// Add more test scenarios here as needed.
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			students, err := repo.GetAll(tc.Page, tc.PageSize)

			assert.NoError(t, err, "Error calling GetAll")
			assert.Equal(t, tc.Expected, students)
		})
	}
}
func TestDelete(t *testing.T) {
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/teststudent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo, err := NewStudentRepository(db)
	if err != nil {
		t.Fatalf("Failed to create student repository: %v", err)
	}

	t.Run("Delete Success", func(t *testing.T) {
		//first add a student with a predefined uuid
		studentID := uuid.New()
		freshStudent := models.StudentEntity{
			ID:      studentID,
			Name:    "Johnny",
			Surname: "Bravo",
		}

		newStudent := EntityToModel(&freshStudent)
		repo.Add(newStudent)

		err = repo.Delete(freshStudent.ID)

		assert.NoError(t, err)
	})

	t.Run("UUID does not exist", func(t *testing.T) {
		err = repo.Delete(uuid.New())
		assert.Nil(t, err)
	})

}

func TestCleanup(t *testing.T) {
	dsn := "kiyam:password@tcp(127.0.0.1:3306)/teststudent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Cleanup: Delete all records from the students table using plain SQL
	if err := db.Exec("DELETE FROM students").Error; err != nil {
		t.Fatalf("Failed to delete records: %v", err)
	}
}

// single page ->> add info
//				     ->all records   -delete
//
