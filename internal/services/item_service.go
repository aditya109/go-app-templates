package services

import (
	model "github.com/aditya109/go-server-template/internal/models"
)

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
