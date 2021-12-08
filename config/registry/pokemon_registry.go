package registry

import (
	"github.com/FernandoGal25/academy-go-q42021/application/usecase"
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"
	"github.com/FernandoGal25/academy-go-q42021/interface/repository"
)

// NewPokemonController instances pokemon controller by inyecting pokemon use case.
func (r registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonService())
}

// NewPokemonService instances pokemon use case by inyecting pokemon repository.
func (r registry) NewPokemonService() usecase.PokemonService {
	return usecase.NewPokemonService(r.NewCSVPokemonRepository(), r.NewRESTPokemonRepository())
}

// NewCSVPokemonRepository instances pokemon repository by inyecting CSV file handler.
func (r registry) NewCSVPokemonRepository() repository.CSVPokemonRepository {
	return repository.NewCSVPokemonRepository(r.csv)
}

// NewRESTPokemonRepository instances pokemon repository by inyecting an api path.
func (r registry) NewRESTPokemonRepository() repository.RESTPokemonRepository {
	return repository.NewRestPokemonRepository(r.api)
}
