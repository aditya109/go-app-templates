package services

import (
	"fmt"

	mwd "github.com/aditya109/go-server-template/internal/middlewares/errors"
	model "github.com/aditya109/go-server-template/internal/models"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

var logger = log.NewLogger()
var items = []model.Item{
	{
		ID:   66,
		Name: "Heidt",
	},
	{
		ID:   12,
		Name: "Bertine",
	},
	{
		ID:   59,
		Name: "Vastah",
	},
	{
		ID:   39,
		Name: "Frendel",
	},
}

func GetAllItems() ([]model.Item, error) {
	return items, nil
}

func GetItemByID(id int64) (model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.ID == id {
			return v, nil
		}
	}
	return model.Item{}, &mwd.AppError{
		Cause: fmt.Sprintf("item with ID=%d not present", id),
	}
}

func GetItemsByIDAndName(id int64, name string) ([]model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.ID == id && v.Name == name {
			return []model.Item{v}, nil
		}
	}
	return []model.Item{}, &mwd.AppError{
		Cause: fmt.Sprintf("item with ID=%d not present", id),
	}
}
