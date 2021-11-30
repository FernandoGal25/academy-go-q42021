package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
)

/*
	Repository that handles the operations of pokemon in a external REST API.
*/
type RESTPokemonRepository struct {
	APIGateway string
}

func NewRestPokemonRepository(p string) RESTPokemonRepository {
	return RESTPokemonRepository{p}
}

func (r RESTPokemonRepository) FindByID(ID uint64) (*model.Pokemon, error) {
	resp, err := http.Get(r.APIGateway + "/pokemon/" + strconv.FormatUint(ID, 10))
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

func (r RESTPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	resp, err := http.Get(r.APIGateway + "/pokemon")
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

func (r RESTPokemonRepository) Persist(*model.Pokemon) error {
	return errors.New("METHOD NOT ALLOWED")
}
