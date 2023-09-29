// Code generated by MockGen. DO NOT EDIT.
// Source: internal/student/services/service.go

// Package services is a generated GoMock package.
package mocks

import (
	models "backend/internal/student/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepository) Add(student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepositoryMockRecorder) Add(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), student)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockRepository) Get(id uuid.UUID) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(page, pageSize int) ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", page, pageSize)
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), page, pageSize)
}

// TotalStudentCount mocks base method.
func (m *MockRepository) TotalStudentCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalStudentCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalStudentCount indicates an expected call of TotalStudentCount.
func (mr *MockRepositoryMockRecorder) TotalStudentCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalStudentCount", reflect.TypeOf((*MockRepository)(nil).TotalStudentCount))
}