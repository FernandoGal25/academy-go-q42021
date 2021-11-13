package registry

import (
	"academy_bootcamp/infrastructure/datastore"
	"academy_bootcamp/interface/controller"
)

type registry struct {
	csv *datastore.CSVHandler
	// api
}

type Registry interface {
	Register() controller.AppController
}

func NewRegistry(csv *datastore.CSVHandler) Registry {
	return &registry{csv}
}

func (r *registry) Register() controller.AppController {
	return controller.AppController{
		Item: r.NewItemController(),
	}
}
