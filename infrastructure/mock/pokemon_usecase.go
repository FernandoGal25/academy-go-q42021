// Code generated by MockGen. DO NOT EDIT.
// Source: application/usecase/pokemon.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/FernandoGal25/academy-go-q42021/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPokemonUsecase is a mock of PokemonUsecase interface.
type MockPokemonUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonUsecaseMockRecorder
}

// MockPokemonUsecaseMockRecorder is the mock recorder for MockPokemonUsecase.
type MockPokemonUsecaseMockRecorder struct {
	mock *MockPokemonUsecase
}

// NewMockPokemonUsecase creates a new mock instance.
func NewMockPokemonUsecase(ctrl *gomock.Controller) *MockPokemonUsecase {
	mock := &MockPokemonUsecase{ctrl: ctrl}
	mock.recorder = &MockPokemonUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokemonUsecase) EXPECT() *MockPokemonUsecaseMockRecorder {
	return m.recorder
}

// CreatePokemon mocks base method.
func (m *MockPokemonUsecase) CreatePokemon(key int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePokemon", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePokemon indicates an expected call of CreatePokemon.
func (mr *MockPokemonUsecaseMockRecorder) CreatePokemon(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePokemon", reflect.TypeOf((*MockPokemonUsecase)(nil).CreatePokemon), key)
}

// GetAllPokemons mocks base method.
func (m *MockPokemonUsecase) GetAllPokemons() ([]model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPokemons")
	ret0, _ := ret[0].([]model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPokemons indicates an expected call of GetAllPokemons.
func (mr *MockPokemonUsecaseMockRecorder) GetAllPokemons() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPokemons", reflect.TypeOf((*MockPokemonUsecase)(nil).GetAllPokemons))
}

// GetPokemonByID mocks base method.
func (m *MockPokemonUsecase) GetPokemonByID(key int) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemonByID", key)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPokemonByID indicates an expected call of GetPokemonByID.
func (mr *MockPokemonUsecaseMockRecorder) GetPokemonByID(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemonByID", reflect.TypeOf((*MockPokemonUsecase)(nil).GetPokemonByID), key)
}

// GetPokemonsByFilters mocks base method.
func (m *MockPokemonUsecase) GetPokemonsByFilters(f map[string]interface{}) ([]model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemonsByFilters", f)
	ret0, _ := ret[0].([]model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPokemonsByFilters indicates an expected call of GetPokemonsByFilters.
func (mr *MockPokemonUsecaseMockRecorder) GetPokemonsByFilters(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemonsByFilters", reflect.TypeOf((*MockPokemonUsecase)(nil).GetPokemonsByFilters), f)
}
