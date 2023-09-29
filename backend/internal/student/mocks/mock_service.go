// Code generated by MockGen. DO NOT EDIT.
// Source: internal/student/controllers/controller.go

// Package controllers is a generated GoMock package.
package mocks

import (
	models "backend/internal/student/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStudentService is a mock of StudentService interface.
type MockStudentService struct {
	ctrl     *gomock.Controller
	recorder *MockStudentServiceMockRecorder
}

// MockStudentServiceMockRecorder is the mock recorder for MockStudentService.
type MockStudentServiceMockRecorder struct {
	mock *MockStudentService
}

// NewMockStudentService creates a new mock instance.
func NewMockStudentService(ctrl *gomock.Controller) *MockStudentService {
	mock := &MockStudentService{ctrl: ctrl}
	mock.recorder = &MockStudentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentService) EXPECT() *MockStudentServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStudentService) Add(student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStudentServiceMockRecorder) Add(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStudentService)(nil).Add), student)
}

// Delete mocks base method.
func (m *MockStudentService) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStudentServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStudentService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockStudentService) Get(id uuid.UUID) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStudentServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStudentService)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockStudentService) GetAll(page, pageSize int) (models.PaginationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", page, pageSize)
	ret0, _ := ret[0].(models.PaginationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStudentServiceMockRecorder) GetAll(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStudentService)(nil).GetAll), page, pageSize)
}