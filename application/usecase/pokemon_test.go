package usecase

import (
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	"github.com/FernandoGal25/academy-go-q42021/interface/repository"

	"github.com/golang/mock/gomock"
)

func TestPokemonService_GetAllPokemons_Correct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := repository.NewMockPokemonRepository(mockCtrl)
	repo.EXPECT().FetchAll().Return([]model.Pokemon{{
		ID:   25,
		Name: "Pikachu",
	}, {
		ID:   249,
		Name: "Lugia",
	}}, nil)
	service := PokemonService{CSVPokemonRepository: repo}
	p, _ := service.GetAllPokemons()
	if p == nil {
		t.Errorf("Pokemons not matching")
	}
}

func TestPokemonService_GetPokemonByID_Correct(t *testing.T) {
	var key uint64 = 25
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := repository.NewMockPokemonRepository(mockCtrl)
	repo.EXPECT().FindByID(key).Return(&model.Pokemon{
		ID:   25,
		Name: "Pikachu",
	}, nil)
	service := PokemonService{CSVPokemonRepository: repo}
	p, _ := service.GetPokemonByID(key)
	if p == nil {
		t.Errorf("Pokemons not matching")
	}
}
