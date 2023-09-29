package services

import (
	"backend/internal/student/mocks"
	"backend/internal/student/models"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ModelToEntity(student *models.Student) *models.StudentEntity {
	return &models.StudentEntity{
		ID:      uuid.MustParse(student.ID),
		Name:    student.Name,
		Surname: student.Surname,
	}
}
func TestGet(t *testing.T) {
	//Get(uuid) --> returns *models.Student
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := Service(repo)

	t.Run("Get Success", func(t *testing.T) {

		expectedStudent := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "hasan",
			Surname: "huseyin",
		}

		student := ModelToEntity(expectedStudent)
		repo.EXPECT().Get(student.ID).Return(expectedStudent, nil)
		actual, err := service.Get(student.ID)
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent, actual)
	})

	t.Run("Get Fail", func(t *testing.T) {

		wrongID := &models.Student{
			ID:      "9999c99f-9d99-9999-9b9f-999d9d9aae99",
			Name:    "hasan",
			Surname: "huseyin",
		}
		nilStudent := ModelToEntity(wrongID)

		expectedError := errors.New("Failed to fetch student")
		repo.EXPECT().Get(nilStudent.ID).Return(nil, expectedError)
		actual, err := service.Get(nilStudent.ID)
		assert.Error(t, err)
		assert.Nil(t, actual)
		assert.Equal(t, expectedError, err)
	})

}

func TestGetAll(t *testing.T) {
	//can't be nil, has default page, pagesize.
	//GetAll(page int, pageSize int) --> returns models.PaginationResponse
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := Service(repo)

	totalStudents := int64(3)
	repo.EXPECT().TotalStudentCount().Return(totalStudents, nil)

	t.Run("GetAll Success", func(t *testing.T) {
		expectedStudents := []models.Student{
			{
				ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
				Name:    "hasan",
				Surname: "huseyin",
			},
			{
				ID:      "1c4f0e9f-5a66-493d-84d4-400e7a7175a1",
				Name:    "ahmet",
				Surname: "talha",
			},
			{
				ID:      "6a6dbce8-ca2a-4473-ae71-b342d7b13545",
				Name:    "matrak",
				Surname: "efe",
			},
		}
		expectedResponse := models.PaginationResponse{
			Students: expectedStudents,
			Page: models.Page{
				Number:   1,
				Size:     10,
				Elements: int(totalStudents),
				Pages:    1,
			},
		}

		page := 1
		pageSize := 10

		repo.EXPECT().GetAll(page, pageSize).Return(expectedStudents, nil)

		actual, err := service.GetAll(page, pageSize)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, actual)
	})
	t.Run("GetAll Fail", func(t *testing.T) {
		page := 0
		pageSize := 0
		expectedError := errors.New("page and pagesize cannot be lower than 1")

		repo.EXPECT().GetAll(page, pageSize).Return(nil, expectedError).Times(0)
		response, err := service.GetAll(page, pageSize)
		nilResponse := models.PaginationResponse{Students: []models.Student(nil), Page: models.Page{Number: 0, Size: 0, Elements: 0, Pages: 0}}

		assert.Equal(t, response, nilResponse)
		assert.Error(t, err)
		assert.Equal(t, err, expectedError)
	})
}

func TestAdd(t *testing.T) {
	//Add(student) --> returns nil
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := Service(repo)
	t.Run("Add Success", func(t *testing.T) {
		student := &models.Student{
			Name:    "keloglan",
			Surname: "kelesoglan",
		}

		repo.EXPECT().Add(gomock.Any()).Return(nil).Times(1)
		err := service.Add(student)

		assert.NoError(t, err)
	})

	t.Run("Add Fail", func(t *testing.T) {
		nilStudent := &models.Student{
			Name:    "",
			Surname: "",
		}

		expectedError := errors.New("name and surname are required")
		repo.EXPECT().Add(nilStudent).Return(expectedError).Times(0)
		err := service.Add(nilStudent)

		assert.Error(t, err)
		assert.Equal(t, err, expectedError)
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	service := Service(repo)

	t.Run("DeleteSuccess", func(t *testing.T) {
		student := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "hasan",
			Surname: "huseyin",
		}

		mockStudent := ModelToEntity(student)
		repo.EXPECT().Delete(mockStudent.ID).Return(nil).Times(1)
		err := service.Delete(mockStudent.ID)

		assert.NoError(t, err)
	})

	t.Run("DeleteFail", func(t *testing.T) {
		student := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "hasan",
			Surname: "huseyin",
		}

		mockStudent := ModelToEntity(student)
		expectedErrorMessage := "expected delete error"
		expectedError := errors.New(expectedErrorMessage)
		repo.EXPECT().Delete(mockStudent.ID).Return(expectedError).Times(1)
		err := service.Delete(mockStudent.ID)

		assert.Error(t, err)
		// t.Log("Error:", err.Error())
		assert.EqualError(t, err, expectedErrorMessage)
	})
}
