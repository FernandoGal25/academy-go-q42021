package registry

import (
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/datastore"
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"
)

type registry struct {
	csv *datastore.CSVHandler
	api string
}

// Registry is the dependency inyector.
type Registry interface {
	Register() controller.AppController
}

// NewRegistry creates new registry.
func NewRegistry(csv *datastore.CSVHandler, api string) registry {
	return registry{csv, api}
}

// Register initializes the dependency inyection.
func (r registry) Register() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
