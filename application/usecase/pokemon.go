package usecase

import (
	"fmt"

	"github.com/FernandoGal25/academy-go-q42021/application/contracts"
	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	errors "github.com/FernandoGal25/academy-go-q42021/error"
)

const totalPokemon = 898

// Service that handles the pokemon case uses.
type PokemonService struct {
	CSVPokemonRepository  contracts.PokemonRepository
	RestPokemonRepository contracts.PokemonRepository
}

// Returns an instance of PokemonService.
func NewPokemonService(r1 contracts.PokemonRepository, r2 contracts.PokemonRepository) PokemonService {
	return PokemonService{r1, r2}
}

func validateRangeID(key int) error {
	if key > totalPokemon {
		return errors.ErrDomainValidation{
			Message: fmt.Sprintf("Out of bound, pokemon with ID: %v searched while there are currently only %v pokemon.", key, totalPokemon),
		}
	} else if key < 1 {
		return errors.ErrDomainValidation{
			Message: fmt.Sprintf("Invalid ID: %v, must be in range between 1 and %v.", key, totalPokemon),
		}
	}

	return nil
}

// Orchestrates the tools required to retrieve a pokemon by ID.
func (s PokemonService) GetPokemonByID(key int) (*model.Pokemon, error) {
	if err := validateRangeID(key); err != nil {
		return nil, err
	}

	p, err := s.CSVPokemonRepository.FindByID(key)

	if err != nil {
		return nil, errors.ErrRepositoryWrapper{Message: "Failed to fetch pokemon from repository", Err: err}
	}

	return p, nil
}

// Orchestrates the tools required to retrieve all pokemon.
func (s PokemonService) GetAllPokemons() ([]model.Pokemon, error) {
	p, err := s.CSVPokemonRepository.FetchAll()

	if err != nil {
		return nil, errors.ErrRepositoryWrapper{Message: "Failed to fetch pokemons from repository", Err: err}
	}

	return p, nil
}

// Orchestrates the tools required to register a new pokemon from
// an external API.
func (s PokemonService) CreatePokemon(key int) (string, error) {
	if err := validateRangeID(key); err != nil {
		return "", err
	}

	p, err := s.RestPokemonRepository.FindByID(key)

	if err != nil {
		return "", errors.ErrRepositoryWrapper{Message: "Failed to fetch pokemon from repository", Err: err}
	}

	if err = s.CSVPokemonRepository.Persist(p); err != nil {
		return "", errors.ErrRepositoryWrapper{Message: "Failed to register pokemon in repository", Err: err}
	}

	return p.Name, nil
}
