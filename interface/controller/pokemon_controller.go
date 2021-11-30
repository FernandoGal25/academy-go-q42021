package controller

import (
	"net/http"

	"github.com/FernandoGal25/academy-go-q42021/application/usecase"
)

/*
	API Gateway, handles the request and response of the
	pokemon endpoints
*/
type PokemonController struct {
	Usecase usecase.PokemonService
}

/*
	Returns an instance of PokemonController
*/
func NewPokemonController(us usecase.PokemonService) PokemonController {
	return PokemonController{us}
}

/*
	Calls GetPokemonByID usecase.
*/
func (ic PokemonController) ActionGetByID(c Context) error {
	key, err := parseIDParam(c)
	if err != nil {
		return responseErrorJSON(c, err)
	}

	result, err := ic.Usecase.GetPokemonByID(key)
	if err != nil {
		return responseErrorJSON(c, err)
	}

	return responseJSON(c, result, http.StatusOK)
}

/*
	Calls GetAllPokemon usecase.
*/
func (ic PokemonController) ActionGetAll(c Context) error {
	result, err := ic.Usecase.GetAllPokemons()
	if err != nil {
		return responseErrorJSON(c, err)
	}

	return responseJSON(c, result, http.StatusOK)
}
