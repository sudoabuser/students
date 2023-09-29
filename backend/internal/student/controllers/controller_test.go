package controllers

import (
	"backend/internal/student/mocks"
	"backend/internal/student/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentService(ctrl)
	controller := &StudentController{
		Service: mockService,
	}

	router := gin.Default()
	router.GET("/students/:id", controller.Get)

	t.Run("GetSuccess", func(t *testing.T) {
		expectedStudent := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "hasan",
			Surname: "huseyin",
		}
		id := uuid.MustParse(expectedStudent.ID)

		mockService.EXPECT().Get(id).Return(expectedStudent, nil)

		w := performRequest(router, "GET", "/students/"+id.String(), nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var actualStudent models.Student
		err := json.Unmarshal(w.Body.Bytes(), &actualStudent)
		assert.NoError(t, err)
		assert.Equal(t, expectedStudent, &actualStudent)
	})

	t.Run("GetFail", func(t *testing.T) {
		id := uuid.New()

		mockService.EXPECT().Get(id).Return(nil, errors.New("student not found"))

		w := performRequest(router, "GET", "/students/"+id.String(), nil)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "student not found", response["message"])
		assert.Equal(t, id.String(), response["student_id"])
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentService(ctrl)
	controller := &StudentController{
		Service: mockService,
	}

	router := gin.Default()
	router.DELETE("/students/:id", controller.Delete)

	t.Run("valid student ID", func(t *testing.T) {
		validID := uuid.New()

		mockService.EXPECT().Delete(validID).Return(nil)

		w := performRequest(router, "DELETE", "/students/"+validID.String(), nil)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Student deleted successfully"}`, w.Body.String())
	})

	t.Run("invalid ID", func(t *testing.T) {
		invalidID := "invalid-uuid"

		w := performRequest(router, "DELETE", "/students/"+invalidID, nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "invalid UUID"}`, w.Body.String())
	})

	t.Run("student not found", func(t *testing.T) {
		notFoundID := uuid.New()

		expectedError := errors.New("student not found")
		mockService.EXPECT().Delete(notFoundID).Return(expectedError)

		w := performRequest(router, "DELETE", "/students/"+notFoundID.String(), nil)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error": "student not found"}`, w.Body.String())
	})
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentService(ctrl)
	controller := &StudentController{
		Service: mockService,
	}

	router := gin.Default()
	router.POST("/students", controller.Add)

	t.Run("ValidStudent", func(t *testing.T) {
		studentToAdd := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "John",
			Surname: "Doe",
		}

		requestBody, _ := json.Marshal(studentToAdd)

		mockService.EXPECT().Add(gomock.Any()).Return(nil)

		w := performRequest(router, "POST", "/students", requestBody)

		assert.Equal(t, http.StatusCreated, w.Code)

		var addedStudent models.Student
		err := json.Unmarshal(w.Body.Bytes(), &addedStudent)
		assert.NoError(t, err)
		assert.Equal(t, studentToAdd, &addedStudent)
	})

	t.Run("InvalidRequestBody", func(t *testing.T) {
		invalidRequestBody := []byte(`{"invalid_json": }`)

		w := performRequest(router, "POST", "/students", invalidRequestBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "invalid request"}`, w.Body.String())
	})

	t.Run("AddError", func(t *testing.T) {
		studentToAdd := &models.Student{
			ID:      "7995c72f-7d04-4136-8b5f-000d6d4aae23",
			Name:    "John",
			Surname: "Doe",
		}

		requestBody, _ := json.Marshal(studentToAdd)

		mockService.EXPECT().Add(gomock.Any()).Return(errors.New("failed to create student"))

		w := performRequest(router, "POST", "/students", requestBody)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "failed to create student"}`, w.Body.String())
	})

}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentService(ctrl)
	controller := &StudentController{
		Service: mockService,
	}

	router := gin.Default()
	router.GET("/students", controller.GetAll)

	t.Run("GetAllSuccess", func(t *testing.T) {
		mockResponse := models.PaginationResponse{
			Students: []models.Student{
				{ID: "1", Name: "John", Surname: "Doe"},
				{ID: "2", Name: "Jane", Surname: "Smith"},
			},
			Page: models.Page{
				Number:   1,
				Size:     10,
				Elements: 2,
				Pages:    1,
			},
		}

		mockService.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(mockResponse, nil)

		w := performRequest(router, "GET", "/students", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var actualResponse models.PaginationResponse
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, mockResponse, actualResponse)
	})

	t.Run("GetAllError", func(t *testing.T) {
		mockService.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(models.PaginationResponse{}, errors.New("failed to retrieve students"))

		w := performRequest(router, "GET", "/students", nil)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "failed to retreive students", response["message"])
	})
}

func performRequest(router *gin.Engine, method, url string, body []byte) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	return w
}
