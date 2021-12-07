package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
)

// Repository that handles the operations of pokemon in a external REST API.
type RESTPokemonRepository struct {
	APIGateway string
}

// Returns an instance of RESTPokemoRepository
func NewRestPokemonRepository(p string) RESTPokemonRepository {
	return RESTPokemonRepository{p}
}

func getHTTPRequest(endpoint string, i interface{}) error {
	resp, err := http.Get(endpoint)

	if err != nil {
		return customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed request at %v", endpoint), Err: err}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Cannot read HTTP response body from %v", endpoint), Err: err}
	}

	if err := json.Unmarshal(body, i); err != nil {
		return customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed to parse json body from HTTP response: %v", endpoint), Err: err}
	}

	return nil
}

// Searches on pokeapi for a pokemon with the given ID.
func (r RESTPokemonRepository) FindByID(ID int) (*model.Pokemon, error) {
	var p model.Pokemon

	if err := getHTTPRequest(fmt.Sprintf("%v/pokemon/%v", r.APIGateway, ID), &p); err != nil {
		return nil, err
	}

	return &p, nil
}

// Fetches all pokemon from pokeapi.
func (r RESTPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	var collection = []model.Pokemon{}

	if err := getHTTPRequest(fmt.Sprintf("%v/pokemon", r.APIGateway), &collection); err != nil {
		return nil, err
	}

	return collection, nil
}

// This method was set in order to comply with contracts.PokemonRepository
// NOTE: Consider creating 2 different contracts in order to delete this.
func (r RESTPokemonRepository) Persist(*model.Pokemon) error {
	return errors.New("method Persist from RestPokemonRepository stil under construction")
}

func (r RESTPokemonRepository) FetchConcurrently(f map[string]interface{}) ([]model.Pokemon, error) {
	return nil, errors.New("method Persist from RestPokemonRepository stil under construction")
}
