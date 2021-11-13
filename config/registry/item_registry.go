package registry

import (
	"academy_bootcamp/application/contracts"
	"academy_bootcamp/application/usecase"
	"academy_bootcamp/interface/controller"
	"academy_bootcamp/interface/repository"
)

func (r *registry) NewItemController() controller.ItemController {
	return controller.NewItemController(r.NewItemService())
}

func (r *registry) NewItemService() usecase.ItemService {
	return usecase.NewItemService(r.NewItemRepository())
}

func (r *registry) NewItemRepository() contracts.ItemRepository {
	return repository.NewItemRepository(r.csv)
}
