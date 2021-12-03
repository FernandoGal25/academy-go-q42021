package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
)

// Repository that handles the operations of pokemon in a external REST API.
type RESTPokemonRepository struct {
	APIGateway string
}

// Returns an instance of RESTPokemoRepository
func NewRestPokemonRepository(p string) RESTPokemonRepository {
	return RESTPokemonRepository{p}
}

// Searches on pokeapi for a pokemon with the given ID.
func (r RESTPokemonRepository) FindByID(ID int) (*model.Pokemon, error) {
	resp, err := http.Get(r.APIGateway + "/pokemon/" + strconv.Itoa(ID))
	if err != nil {
		return nil, errors.New("REQUEST ERROR")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("FORMAT ERROR")
	}

	var p model.Pokemon
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, errors.New(string(body))
	}

	return &p, nil
}

// Fetches all pokemon from pokeapi.
func (r RESTPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	resp, err := http.Get(fmt.Sprintf(r.APIGateway + "/pokemon"))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var collection = []model.Pokemon{}

	if err := json.Unmarshal(body, &collection); err != nil {
		return nil, err
	}

	return collection, nil
}

// This method was set in order to comply with contracts.PokemonRepository
// NOTE: Consider creating 2 different contracts in order to delete this.
func (r RESTPokemonRepository) Persist(*model.Pokemon) error {
	return errors.New("METHOD NOT ALLOWED")
}
