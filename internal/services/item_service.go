package services

import (
	"fmt"
	"time"

	mwd "github.com/aditya109/go-server-template/internal/middlewares/errors"
	model "github.com/aditya109/go-server-template/internal/models"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

var logger = log.NewLogger()

func GetAllItems() ([]model.Item, error) {
	var items = []model.Item{
		{
			Id:   66,
			Name: "Heidt",
		},
		{
			Id:   12,
			Name: "Bertine",
		},
		{
			Id:   59,
			Name: "Vastah",
		},
		{
			Id:   39,
			Name: "Frendel",
		},
	}
	return items, nil
}

func GetItemById(id int64) (model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.Id == id {
			return v, nil
		}
	}
	return model.Item{}, &mwd.AppError{
		When:  time.Now(),
		Cause: fmt.Sprintf("item with Id=%d not present", id),
	}
}

func GetItemsByIdAndName(id int64, name string) ([]model.Item, error) {
	items, err := GetAllItems()
	if err != nil {
		logger.Error(err)
	}
	for _, v := range items {
		if v.Id == id && v.Name == name {
			return []model.Item{v}, nil
		}
	}
	return []model.Item{}, &mwd.AppError{
		When:  time.Now(),
		Cause: fmt.Sprintf("item with Id=%d not present", id),
	}
}
