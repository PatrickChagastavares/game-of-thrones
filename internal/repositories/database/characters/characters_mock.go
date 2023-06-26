// Code generated by MockGen. DO NOT EDIT.
// Source: characters.go

// Package characters is a generated GoMock package.
package characters

import (
	context "context"
	reflect "reflect"

	entities "github.com/PatrickChagastavares/game-of-thrones/internal/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRepository) Create(ctx context.Context, character entities.CharacterRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, character)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRepositoryMockRecorder) Create(ctx, character interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRepository)(nil).Create), ctx, character)
}

// Delete mocks base method.
func (m *MockIRepository) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRepository)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockIRepository) Find(ctx context.Context) ([]entities.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx)
	ret0, _ := ret[0].([]entities.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIRepositoryMockRecorder) Find(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIRepository)(nil).Find), ctx)
}

// FindByID mocks base method.
func (m *MockIRepository) FindByID(ctx context.Context, id string) (entities.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(entities.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockIRepositoryMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIRepository)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockIRepository) Update(ctx context.Context, character *entities.Character) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, character)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRepositoryMockRecorder) Update(ctx, character interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRepository)(nil).Update), ctx, character)
}