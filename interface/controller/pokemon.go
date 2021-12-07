package controller

import (
	"net/http"
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/application/usecase"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
)

// PokemonController handles the request and response of the pokemon endpoints
type PokemonController struct {
	Usecase usecase.PokemonUsecase
}

// NewPokemonController returns an instance of PokemonController.
func NewPokemonController(us usecase.PokemonUsecase) PokemonController {
	return PokemonController{us}
}

// ActionGetByID calls GetPokemonByID usecase.
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

// ActionGetAll calls GetAllPokemon usecase.
func (ic PokemonController) ActionGetAll(c Context) error {
	result, err := ic.Usecase.GetAllPokemons()
	if err != nil {
		return responseErrorJSON(c, err)
	}

	return responseJSON(c, result, http.StatusOK)
}

// ActionPostByID calls CreatePokemon usecase.
func (ic PokemonController) ActionPostByID(c Context) error {
	key, err := parseIDParam(c)
	if err != nil {
		return responseErrorJSON(c, err)
	}

	name, err := ic.Usecase.CreatePokemon(key)
	if err != nil {
		return responseErrorJSON(c, err)
	}

	return responseJSON(
		c,
		map[string]string{"Message": name + " has been registered in the pokedex."},
		http.StatusCreated,
	)
}

func adaptFilters(c Context) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	qp := c.QueryParams()

	if qp["type"][0] == "odd" {
		filters["id"] = func(id int) bool {
			return id%2 == 1
		}
	} else if qp["type"][0] == "even" {
		filters["id"] = func(id int) bool {
			return id%2 == 0
		}
	} else {
		filters["id"] = func(id int) bool {
			return true
		}
	}

	if qp["items"] != nil {
		limit, err := strconv.Atoi(qp["items"][0])

		if err != nil {
			return nil, customErrors.ErrInvalidRequest{Message: "Invalid 'items' queryParam", Err: err}
		}

		filters["limit"] = limit
	} else {
		filters["limit"] = 10
	}

	if qp["items_per_workers"] != nil {
		wj, err := strconv.Atoi(qp["items_per_workers"][0])

		if err != nil {
			return nil, customErrors.ErrInvalidRequest{Message: "Invalid 'items_per_workers' queryParam", Err: err}
		}
		filters["workerJobs"] = wj
	} else {
		filters["workerJobs"] = 4
	}

	return filters, nil
}

// ActionGetByFilters calls GetPokemonsByFilters.
func (ic PokemonController) ActionGetByFilters(c Context) error {
	filters, err := adaptFilters(c)

	if err != nil {
		return responseErrorJSON(c, err)
	}

	result, err := ic.Usecase.GetPokemonsByFilters(filters)

	if err != nil {
		return responseErrorJSON(c, err)
	}

	return responseJSON(c, result, http.StatusOK)
}
