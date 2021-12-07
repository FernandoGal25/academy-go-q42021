package controller

import (
	"strconv"
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPokemonController_ActionGetAll(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(usecase *mock.MockPokemonUsecase, context *mock.MockContext)
		wantErr error
	}{
		{
			name: "Correct test",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().GetAllPokemons().Return([]model.Pokemon{{
					ID:             4,
					Name:           "charmander",
					Height:         6,
					Weight:         85,
					Order:          5,
					BaseExperience: 62,
				}, {
					ID:             5,
					Name:           "charmeleon",
					Height:         11,
					Weight:         190,
					Order:          6,
					BaseExperience: 142,
				}, {
					ID:             6,
					Name:           "charizard",
					Height:         17,
					Weight:         905,
					Order:          7,
					BaseExperience: 240,
				}}, nil)
				context.EXPECT().JSON(200, []model.Pokemon{{
					ID:             4,
					Name:           "charmander",
					Height:         6,
					Weight:         85,
					Order:          5,
					BaseExperience: 62,
				}, {
					ID:             5,
					Name:           "charmeleon",
					Height:         11,
					Weight:         190,
					Order:          6,
					BaseExperience: 142,
				}, {
					ID:             6,
					Name:           "charizard",
					Height:         17,
					Weight:         905,
					Order:          7,
					BaseExperience: 240,
				}}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "Usecase error",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().GetAllPokemons().Return(nil, assert.AnError)
				context.EXPECT().JSON(500, map[string][]ErrorResponse{"errors": {
					{Message: assert.AnError.Error(), ErrorType: "*errors.errorString"},
				}}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	usecase := mock.NewMockPokemonUsecase(mockCtrl)
	ctrl := NewPokemonController(usecase)

	for _, tt := range tests {
		context := mock.NewMockContext(mockCtrl)
		if tt.prepare != nil {
			tt.prepare(usecase, context)
		}
		err := ctrl.ActionGetAll(context)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}

func TestPokemonController_ActionGetByID(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(usecase *mock.MockPokemonUsecase, context *mock.MockContext)
		wantErr error
	}{
		{
			name: "Correct test",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().GetPokemonByID(7).Return(&model.Pokemon{
					ID:             7,
					Name:           "squirtle",
					Height:         5,
					Weight:         90,
					Order:          10,
					BaseExperience: 63,
				}, nil)
				context.EXPECT().Param("id").Return("7")
				context.EXPECT().JSON(200, &model.Pokemon{
					ID:             7,
					Name:           "squirtle",
					Height:         5,
					Weight:         90,
					Order:          10,
					BaseExperience: 63,
				}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "Bad ID param",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				_, err := strconv.Atoi("a")
				context.EXPECT().Param("id").Return("a")
				context.EXPECT().JSON(400, map[string][]ErrorResponse{"errors": {
					{Message: "Invalid ID param, must be a number", ErrorType: "errors.ErrInvalidRequest"},
					{Message: err.Error(), ErrorType: "*strconv.NumError"},
					{Message: "invalid syntax", ErrorType: "*errors.errorString"},
				}}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Usecase error",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().GetPokemonByID(1).Return(nil, assert.AnError)
				context.EXPECT().Param("id").Return("1")
				context.EXPECT().JSON(500, map[string][]ErrorResponse{"errors": {
					{Message: assert.AnError.Error(), ErrorType: "*errors.errorString"},
				}}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	usecase := mock.NewMockPokemonUsecase(mockCtrl)
	ctrl := NewPokemonController(usecase)

	for _, tt := range tests {
		context := mock.NewMockContext(mockCtrl)
		if tt.prepare != nil {
			tt.prepare(usecase, context)
		}
		err := ctrl.ActionGetByID(context)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}

func TestPokemonController_ActionPostById(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(usecase *mock.MockPokemonUsecase, context *mock.MockContext)
		wantErr error
	}{
		{
			name: "Correct test",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().CreatePokemon(7).Return("squirtle", nil)
				context.EXPECT().Param("id").Return("7")
				context.EXPECT().JSON(201, map[string]string{"Message": "squirtle has been registered in the pokedex."}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "Bad ID param",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				_, err := strconv.Atoi("a")
				context.EXPECT().Param("id").Return("a")
				context.EXPECT().JSON(400, map[string][]ErrorResponse{"errors": {
					{Message: "Invalid ID param, must be a number", ErrorType: "errors.ErrInvalidRequest"},
					{Message: err.Error(), ErrorType: "*strconv.NumError"},
					{Message: "invalid syntax", ErrorType: "*errors.errorString"},
				}}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Usecase error",
			prepare: func(usecase *mock.MockPokemonUsecase, context *mock.MockContext) {
				usecase.EXPECT().CreatePokemon(1).Return("", assert.AnError)
				context.EXPECT().Param("id").Return("1")
				context.EXPECT().JSON(500, map[string][]ErrorResponse{"errors": {
					{Message: assert.AnError.Error(), ErrorType: "*errors.errorString"},
				}}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	usecase := mock.NewMockPokemonUsecase(mockCtrl)
	ctrl := NewPokemonController(usecase)

	for _, tt := range tests {
		context := mock.NewMockContext(mockCtrl)
		if tt.prepare != nil {
			tt.prepare(usecase, context)
		}
		err := ctrl.ActionPostByID(context)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}
