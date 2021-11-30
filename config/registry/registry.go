package registry

import (
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/datastore"
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"
)

type registry struct {
	csv *datastore.CSVHandler
}

/*
	Dependency inyector.
*/
type Registry interface {
	Register() controller.AppController
}

/*
	Creates new registry.
*/
func NewRegistry(csv *datastore.CSVHandler) registry {
	return registry{csv}
}

/*
	Initializes the dependency inyection.
*/
func (r registry) Register() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
