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
	CSVPokemonRepository  contracts.PokemonRepository
	RestPokemonRepository contracts.PokemonRepository
}

/*
	Returns an instance of PokemonService.
*/
func NewPokemonService(r1 contracts.PokemonRepository, r2 contracts.PokemonRepository) PokemonService {
	return PokemonService{r1, r2}
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
	p, err := s.CSVPokemonRepository.FindByID(key)
	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
	Orchestrates the tools required to retrieve all pokemon.
*/
func (s PokemonService) GetAllPokemons() ([]model.Pokemon, error) {
	p, err := s.CSVPokemonRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
	Orchestrates the tools required to register a new pokemon from
	an external API.
*/
func (s PokemonService) CreatePokemon(key uint64) (*model.Pokemon, error) {
	p, err := s.RestPokemonRepository.FindByID(key)
	if err != nil {
		return nil, err
	}

	return p, s.CSVPokemonRepository.Persist(p)
}
