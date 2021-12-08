package repository

import (
	"strconv"
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCSVPokemonRepository_FindByID(t *testing.T) {
	header := []string{
		"id", "name", "height", "weight", "order", "base_experience",
	}
	readMock := []string{"1", "bulbasaur", "7", "69", "1", "64"}
	readBad := []string{"a", "test fail", "1", "1", "1", "1"}
	_, badFormatErr := strconv.Atoi("a")
	tests := []struct {
		name       string
		key        int
		prepare    func(handler *mock.MockFileManager)
		wantResult *model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			key:  1,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().GetHeader().Return(header)
				handler.EXPECT().Read().Return(readMock, nil)
				handler.EXPECT().Close()
			},
			wantResult: &model.Pokemon{
				ID:             1,
				Name:           "bulbasaur",
				Height:         7,
				Weight:         69,
				Order:          1,
				BaseExperience: 64,
			},
			wantErr: nil,
		},
		{
			name: "CSVHandler.BuildHandler() error",
			key:  1,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrDatastoreWrapper{Message: "Could not initialize CSVHandler", Err: assert.AnError},
		},
		{
			name: "CSV read error",
			key:  1,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().Read().Return(nil, assert.AnError)
				handler.EXPECT().Close()
			},
			wantResult: nil,
			wantErr:    customErrors.ErrDatastoreWrapper{Message: "Failed to read from datastore", Err: assert.AnError},
		},
		{
			name: "CSV bad format",
			key:  1,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().Read().Return(readBad, nil)
				handler.EXPECT().Close()
			},
			wantResult: nil,
			wantErr:    customErrors.ErrCSVFormat{Message: "CSV item does not possess a valid ID", Err: badFormatErr},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		handler := mock.NewMockFileManager(mockCtrl)
		repo := NewCSVPokemonRepository(handler)
		if tt.prepare != nil {
			tt.prepare(handler)
		}
		got, err := repo.FindByID(tt.key)
		assert.Equal(t, tt.wantResult, got, "%v , expected: %v, got: %v", tt.name, tt.wantResult, got)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}

func TestCSVPokemonRepository_FetchAll(t *testing.T) {
	header := []string{
		"id", "name", "height", "weight", "order", "base_experience",
	}
	readMock := [][]string{{"1", "bulbasaur", "7", "69", "1", "64"}}
	readBad := [][]string{{"a", "test fail", "1", "1", "1", "1"}}
	_, badFormatErr := strconv.Atoi("a")
	tests := []struct {
		name       string
		prepare    func(handler *mock.MockFileManager)
		wantResult []model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().GetHeader().Return(header)
				handler.EXPECT().ReadAll().Return(readMock, nil)
				handler.EXPECT().Close()
			},
			wantResult: []model.Pokemon{{
				ID:             1,
				Name:           "bulbasaur",
				Height:         7,
				Weight:         69,
				Order:          1,
				BaseExperience: 64,
			}},
			wantErr: nil,
		},
		{
			name: "CSVHandler.BuildHandler() error",
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrDatastoreWrapper{Message: "Could not initialize CSVHandler", Err: assert.AnError},
		},
		{
			name: "CSV read error",
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().ReadAll().Return(nil, assert.AnError)
				handler.EXPECT().Close()
			},
			wantResult: nil,
			wantErr:    customErrors.ErrDatastoreWrapper{Message: "Failed to fetch all items", Err: assert.AnError},
		},
		{
			name: "CSV bad format",
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().ReadAll().Return(readBad, nil)
				handler.EXPECT().Close()
			},
			wantResult: nil,
			wantErr:    customErrors.ErrCSVFormat{Message: "CSV item does not possess a valid ID", Err: badFormatErr},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		handler := mock.NewMockFileManager(mockCtrl)
		repo := NewCSVPokemonRepository(handler)
		if tt.prepare != nil {
			tt.prepare(handler)
		}
		got, err := repo.FetchAll()
		assert.Equal(t, tt.wantResult, got, "%v , expected: %v, got: %v", tt.name, tt.wantResult, got)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}

func TestCSVPokemonRepository_Persist(t *testing.T) {
	p := model.Pokemon{
		ID:             1,
		Name:           "bulbasaur",
		Height:         7,
		Weight:         69,
		Order:          1,
		BaseExperience: 64,
	}
	w := []string{"1", "bulbasaur", "7", "69", "1", "64"}
	tests := []struct {
		name    string
		model   *model.Pokemon
		prepare func(handler *mock.MockFileManager)
		wantErr error
	}{
		{
			name:  "Correct test",
			model: &p,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().Write(w).Return(nil)
				handler.EXPECT().Close()
			},
			wantErr: nil,
		},
		{
			name:  "CSVHandler.BuildHandler() error",
			model: &p,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(assert.AnError)
			},
			wantErr: customErrors.ErrDatastoreWrapper{Message: "Could not initialize CSVHandler", Err: assert.AnError},
		},
		{
			name:  "CSV write error",
			model: &p,
			prepare: func(handler *mock.MockFileManager) {
				handler.EXPECT().BuildHandler().Return(nil)
				handler.EXPECT().Write(w).Return(assert.AnError)
				handler.EXPECT().Close()
			},
			wantErr: customErrors.ErrDatastoreWrapper{Message: "Failed to persist item", Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		handler := mock.NewMockFileManager(mockCtrl)
		repo := NewCSVPokemonRepository(handler)
		if tt.prepare != nil {
			tt.prepare(handler)
		}
		err := repo.Persist(tt.model)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}
