// Code generated by MockGen. DO NOT EDIT.
// Source: houses.go

// Package houses is a generated GoMock package.
package houses

import (
	context "context"
	reflect "reflect"

	entities "github.com/PatrickChagastavares/game-of-thrones/internal/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIService) Create(ctx context.Context, newHouse entities.HouseRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, newHouse)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIServiceMockRecorder) Create(ctx, newHouse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIService)(nil).Create), ctx, newHouse)
}

// Delete mocks base method.
func (m *MockIService) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIService)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockIService) Find(ctx context.Context, name string) ([]entities.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, name)
	ret0, _ := ret[0].([]entities.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIServiceMockRecorder) Find(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIService)(nil).Find), ctx, name)
}

// FindByID mocks base method.
func (m *MockIService) FindByID(ctx context.Context, id string) (entities.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(entities.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockIServiceMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIService)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockIService) Update(ctx context.Context, updateHouse entities.HouseRequest) (entities.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, updateHouse)
	ret0, _ := ret[0].(entities.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIServiceMockRecorder) Update(ctx, updateHouse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIService)(nil).Update), ctx, updateHouse)
}
