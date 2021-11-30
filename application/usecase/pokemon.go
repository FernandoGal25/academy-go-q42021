package usecase

import (
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/application/contracts"
	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	errors "github.com/FernandoGal25/academy-go-q42021/error"
)

const TOTAL_POKEMON = 898

/*
	Service that handles the pokemon case uses.
*/
type PokemonService struct {
	PokemonRepository contracts.PokemonRepository
}

/*
	Returns an instance of PokemonService.
*/
func NewPokemonService(r contracts.PokemonRepository) PokemonService {
	return PokemonService{r}
}

/*
	Orchestrates the tools required to retrieve a pokemon by ID.
*/
func (s PokemonService) GetPokemonByID(key uint64) (*model.Pokemon, error) {
	if key > TOTAL_POKEMON {
		return nil, errors.DomainValidationError{
			Message: "OUT OF BOUND, CURRENTLY THERE ARE ONLY " + strconv.Itoa(TOTAL_POKEMON) + " POKEMON",
		}
	}
	p, err := s.PokemonRepository.FindByID(key)
	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
	Orchestrates the tools required to retrieve all pokemon.
*/
func (s PokemonService) GetAllPokemons() ([]model.Pokemon, error) {
	p, err := s.PokemonRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return p, nil
}
