// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/repositories.go

// Package mockups is a generated GoMock package.
package mockups

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"
	reflect "reflect"
)

// MockGamesRepository is a mock of GamesRepository interface
type MockGamesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGamesRepositoryMockRecorder
}

// MockGamesRepositoryMockRecorder is the mock recorder for MockGamesRepository
type MockGamesRepositoryMockRecorder struct {
	mock *MockGamesRepository
}

// NewMockGamesRepository creates a new mock instance
func NewMockGamesRepository(ctrl *gomock.Controller) *MockGamesRepository {
	mock := &MockGamesRepository{ctrl: ctrl}
	mock.recorder = &MockGamesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGamesRepository) EXPECT() *MockGamesRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockGamesRepository) Get(id string) (domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGamesRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGamesRepository)(nil).Get), id)
}

// Create mocks base method
func (m *MockGamesRepository) Create(arg0 domain.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGamesRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGamesRepository)(nil).Create), arg0)
}

// Update mocks base method
func (m *MockGamesRepository) Update(arg0 domain.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGamesRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGamesRepository)(nil).Update), arg0)
}