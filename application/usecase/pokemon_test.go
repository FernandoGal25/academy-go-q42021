package usecase

import (
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/FernandoGal25/academy-go-q42021/interface/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPokemonService_GetAllPokemons(t *testing.T) {
	tests := []struct {
		name       string
		prepare    func(repo *repository.MockPokemonRepository)
		wantResult []model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FetchAll().Return([]model.Pokemon{{
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
			},
			wantResult: []model.Pokemon{{
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
			}},
			wantErr: nil,
		},
		{
			name: "Repository error test",
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FetchAll().Return(nil, assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrRepositoryWrapper{Message: "Failed to fetch pokemons from repository", Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := repository.NewMockPokemonRepository(mockCtrl)
	service := PokemonService{CSVPokemonRepository: repo}

	for _, tt := range tests {
		if tt.prepare != nil {
			tt.prepare(repo)
		}
		result, err := service.GetAllPokemons()
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
		assert.Equal(t, tt.wantResult, result, "%v , expected: %v, got: %v", tt.name, tt.wantResult, result)
	}

}

func TestPokemonService_GetPokemonByID(t *testing.T) {

	tests := []struct {
		name       string
		key        int
		prepare    func(repo *repository.MockPokemonRepository)
		wantResult *model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			key:  7,
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FindByID(7).Return(&model.Pokemon{
					ID:             7,
					Name:           "squirtle",
					Height:         5,
					Weight:         90,
					Order:          10,
					BaseExperience: 63,
				}, nil)
			},
			wantResult: &model.Pokemon{
				ID:             7,
				Name:           "squirtle",
				Height:         5,
				Weight:         90,
				Order:          10,
				BaseExperience: 63,
			},
			wantErr: nil,
		},
		{
			name:       "Out of bound ID test",
			key:        900,
			prepare:    nil,
			wantResult: nil,
			wantErr:    customErrors.ErrDomainValidation{Message: "Out of bound, pokemon with ID: 900 searched while there are currently only 898 pokemon.", Err: nil},
		},
		{
			name: "Repository error test",
			key:  300,
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FindByID(300).Return(nil, assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrRepositoryWrapper{Message: "Failed to fetch pokemon from repository", Err: assert.AnError},
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := repository.NewMockPokemonRepository(mockCtrl)
	service := PokemonService{CSVPokemonRepository: repo}
	for _, tt := range tests {

		if tt.prepare != nil {
			tt.prepare(repo)
		}

		result, err := service.GetPokemonByID(tt.key)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
		assert.Equal(t, tt.wantResult, result, "%v , expected: %v, got: %v", tt.name, tt.wantResult, result)
	}
}

func TestPokemonService_CreatePokemon(t *testing.T) {
	tests := []struct {
		name       string
		key        int
		prepare    func(repo1 *repository.MockPokemonRepository, repo2 *repository.MockPokemonRepository)
		wantResult string
		wantErr    error
	}{
		{
			name: "Correct test",
			key:  7,
			prepare: func(repo1 *repository.MockPokemonRepository, repo2 *repository.MockPokemonRepository) {
				gomock.InOrder(repo1.EXPECT().FindByID(7).Return(&model.Pokemon{
					ID:             7,
					Name:           "squirtle",
					Height:         5,
					Weight:         90,
					Order:          10,
					BaseExperience: 63,
				}, nil), repo2.EXPECT().Persist(&model.Pokemon{
					ID:             7,
					Name:           "squirtle",
					Height:         5,
					Weight:         90,
					Order:          10,
					BaseExperience: 63,
				}).Return(nil),
				)
			},
			wantResult: "squirtle",
			wantErr:    nil,
		},
		{
			name:       "Out of bound ID test",
			key:        900,
			prepare:    nil,
			wantResult: "",
			wantErr:    customErrors.ErrDomainValidation{Message: "Out of bound, pokemon with ID: 900 searched while there are currently only 898 pokemon.", Err: nil},
		},
		{
			name: "REST Repository read error test",
			key:  300,
			prepare: func(repo1 *repository.MockPokemonRepository, repo2 *repository.MockPokemonRepository) {
				repo1.EXPECT().FindByID(300).Return(nil, assert.AnError)
			},
			wantResult: "",
			wantErr:    customErrors.ErrRepositoryWrapper{Message: "Failed to fetch pokemon from repository", Err: assert.AnError},
		},
		{
			name: "CSV Repository write error test",
			key:  7,
			prepare: func(repo1 *repository.MockPokemonRepository, repo2 *repository.MockPokemonRepository) {
				gomock.InOrder(
					repo1.EXPECT().FindByID(7).Return(&model.Pokemon{
						ID:             7,
						Name:           "squirtle",
						Height:         5,
						Weight:         90,
						Order:          10,
						BaseExperience: 63,
					}, nil),
					repo2.EXPECT().Persist(&model.Pokemon{
						ID:             7,
						Name:           "squirtle",
						Height:         5,
						Weight:         90,
						Order:          10,
						BaseExperience: 63,
					}).Return(assert.AnError),
				)
			},
			wantResult: "",
			wantErr:    customErrors.ErrRepositoryWrapper{Message: "Failed to register pokemon in repository", Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo1 := repository.NewMockPokemonRepository(mockCtrl)
	repo2 := repository.NewMockPokemonRepository(mockCtrl)
	service := PokemonService{CSVPokemonRepository: repo2, RestPokemonRepository: repo1}
	for _, tt := range tests {

		if tt.prepare != nil {
			tt.prepare(repo1, repo2)
		}

		result, err := service.CreatePokemon(tt.key)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
		assert.Equal(t, tt.wantResult, result, "%v , expected: %v, got: %v", tt.name, tt.wantResult, result)
	}
}

func TestPokemonService_GetPokemonsByFilters(t *testing.T) {
	filters := map[string]interface{}{
		"id": func(id int) bool {
			return id%2 == 1
		},
		"limit":      4,
		"workerJobs": 2,
	}
	tests := []struct {
		name       string
		filters    map[string]interface{}
		prepare    func(repo *repository.MockPokemonRepository)
		wantResult []model.Pokemon
		wantErr    error
	}{
		{
			name:    "Correct test",
			filters: filters,
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FetchConcurrently(filters).Return([]model.Pokemon{{
					ID:             1,
					Name:           "bulbasaur",
					Height:         7,
					Weight:         69,
					Order:          1,
					BaseExperience: 64,
				}, {
					ID:             3,
					Name:           "venusaur",
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
					ID:             7,
					Name:           "squirtle",
					Height:         17,
					Weight:         905,
					Order:          7,
					BaseExperience: 240,
				}}, nil)
			},
			wantResult: []model.Pokemon{{
				ID:             1,
				Name:           "bulbasaur",
				Height:         7,
				Weight:         69,
				Order:          1,
				BaseExperience: 64,
			}, {
				ID:             3,
				Name:           "venusaur",
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
				ID:             7,
				Name:           "squirtle",
				Height:         17,
				Weight:         905,
				Order:          7,
				BaseExperience: 240,
			}},
			wantErr: nil,
		},
		{
			name:    "Repository error test",
			filters: filters,
			prepare: func(repo *repository.MockPokemonRepository) {
				repo.EXPECT().FetchConcurrently(filters).Return(nil, assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrRepositoryWrapper{Message: "Failed to get pokemons with filters from repository", Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := repository.NewMockPokemonRepository(mockCtrl)
	service := PokemonService{CSVPokemonRepository: repo}
	for _, tt := range tests {

		if tt.prepare != nil {
			tt.prepare(repo)
		}

		result, err := service.GetPokemonsByFilters(tt.filters)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
		assert.Equal(t, tt.wantResult, result, "%v , expected: %v, got: %v", tt.name, tt.wantResult, result)
	}
}
