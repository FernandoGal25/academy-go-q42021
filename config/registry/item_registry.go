package registry

import (
	"github.com/FernandoGal25/academy-go-q42021/application/usecase"
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"
	"github.com/FernandoGal25/academy-go-q42021/interface/repository"
)

/*
	Instances pokemon controller by inyecting pokemon use case.
*/
func (r registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonService())
}

/*
	Instances pokemon use case by inyecting pokemon repository.
*/
func (r registry) NewPokemonService() usecase.PokemonService {
	return usecase.NewPokemonService(r.NewCSVPokemonRepository(), r.NewRESTPokemonRepository())
}

/*
	Instances pokemon repository by inyecting CSV file handler.
*/
func (r registry) NewCSVPokemonRepository() repository.CSVPokemonRepository {
	return repository.NewCSVPokemonRepository(r.csv)
}

func (r registry) NewRESTPokemonRepository() repository.RESTPokemonRepository {
	return repository.NewRestPokemonRepository(r.api)
}
