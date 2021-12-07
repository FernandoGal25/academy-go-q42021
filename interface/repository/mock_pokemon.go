// Code generated by MockGen. DO NOT EDIT.
// Source: application/contracts/pokemon_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	model "github.com/FernandoGal25/academy-go-q42021/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPokemonRepository is a mock of PokemonRepository interface.
type MockPokemonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonRepositoryMockRecorder
}

// MockPokemonRepositoryMockRecorder is the mock recorder for MockPokemonRepository.
type MockPokemonRepositoryMockRecorder struct {
	mock *MockPokemonRepository
}

// NewMockPokemonRepository creates a new mock instance.
func NewMockPokemonRepository(ctrl *gomock.Controller) *MockPokemonRepository {
	mock := &MockPokemonRepository{ctrl: ctrl}
	mock.recorder = &MockPokemonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokemonRepository) EXPECT() *MockPokemonRepositoryMockRecorder {
	return m.recorder
}

// FetchAll mocks base method.
func (m *MockPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAll")
	ret0, _ := ret[0].([]model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll.
func (mr *MockPokemonRepositoryMockRecorder) FetchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockPokemonRepository)(nil).FetchAll))
}

// FetchConcurrently mocks base method.
func (m *MockPokemonRepository) FetchConcurrently(f map[string]interface{}) ([]model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchConcurrently", f)
	ret0, _ := ret[0].([]model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchConcurrently indicates an expected call of FetchConcurrently.
func (mr *MockPokemonRepositoryMockRecorder) FetchConcurrently(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchConcurrently", reflect.TypeOf((*MockPokemonRepository)(nil).FetchConcurrently), f)
}

// FindByID mocks base method.
func (m *MockPokemonRepository) FindByID(id int) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPokemonRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPokemonRepository)(nil).FindByID), id)
}

// Persist mocks base method.
func (m *MockPokemonRepository) Persist(arg0 *model.Pokemon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockPokemonRepositoryMockRecorder) Persist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockPokemonRepository)(nil).Persist), arg0)
}
