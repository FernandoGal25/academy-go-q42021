package controller

import (
	"errors"
	"net/http"
	"strconv"

	"academy_bootcamp/application/usecase"
	"academy_bootcamp/domain/model"
)

type itemController struct {
	Usecase usecase.ItemService
}

type ItemController interface {
	GetItemById(c Context) error
	// CreateItem(c Context) error
}

func NewItemController(us usecase.ItemService) ItemController {
	return &itemController{us}
}

func (uc *itemController) GetItemById(c Context) error {
	var item *model.Item
	key, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		return errors.New("QueryParam invalido")
	}

	item, err2 := uc.Usecase.GetByKey(key)

	if err2 != nil {
		return err2
	}
	return c.JSON(http.StatusOK, item)
}

/*func (uc *itemController) CreateItem(c Context) error {
	var params model.Item

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	u, err := uc.u.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
*/
