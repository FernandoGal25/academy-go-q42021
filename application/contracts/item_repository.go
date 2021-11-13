package contracts

import "academy_bootcamp/domain/model"

type ItemRepository interface {
	FindByID(id uint64) (*model.Item, error)
	// Create(u *model.Item) (*model.Item, error)
}
