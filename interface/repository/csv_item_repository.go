package repository

import (
	"academy_bootcamp/application/contracts"
	"academy_bootcamp/domain/model"
	"academy_bootcamp/infrastructure/datastore"
	"errors"
	"io"
	"strconv"
)

type CSVItemRepository struct {
	Handler *datastore.CSVHandler
}

func NewItemRepository(h *datastore.CSVHandler) contracts.ItemRepository {
	return &CSVItemRepository{h}
}

func (r *CSVItemRepository) FindByID(id uint64) (*model.Item, error) {
	var rec *model.Item

	erro := r.Handler.BuildHandler()

	if erro != nil {
		return nil, erro
	}
	for {

		record, err := r.Handler.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, errors.New("no pudo leer el csv")
		}

		key, err2 := strconv.ParseUint(record[0], 10, 16)

		if err2 != nil {
			return nil, errors.New("no es una llave")
		}

		if key == id {
			rec = &model.Item{ID: key, Name: record[1]}
			break
		}
	}
	r.Handler.File.Close()

	return rec, nil
}

/*
func (ur *CSVItemRepository) Create(u *model.Item) (*model.Item, error) {
	if err := ur.csv.Create(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}*/
