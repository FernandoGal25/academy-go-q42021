package usecase

import (
	"academy_bootcamp/application/contracts"
	"academy_bootcamp/domain/model"
)

type itemService struct {
	ItemRepository contracts.ItemRepository
}

type ItemService interface {
	GetByKey(key uint64) (*model.Item, error)
	// Create(u *model.Item) (*model.Item, error)
}

func NewItemService(r contracts.ItemRepository) ItemService {
	return &itemService{r}
}

func (s *itemService) GetByKey(key uint64) (*model.Item, error) {
	u, err := s.ItemRepository.FindByID(key)
	if err != nil {
		return nil, err
	}

	return u, nil
}

/*func (us *itemService) Create(u *model.Item) (*model.Item, error) {
	u, err := us.ItemRepository.Create(u)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}*/
