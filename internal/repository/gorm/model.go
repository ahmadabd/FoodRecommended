package gorm

import (
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

type Food struct {
	Id         int             `json:"id" gorm:"primary_key"`
	Name       string          `json:"name"`
	City       string          `json:"city"`
	Country    string          `json:"country"`
	Vegetarian enum.Vegetative `json:"vegetarian"`
}

func mapFromFoodEntity(food model.Food) Food {
	return Food{
		Id:         food.Id,
		Name:       food.Name,
		City:       food.City,
		Country:    food.Country,
		Vegetarian: food.Vegetarian,
	}
}

func mapToFoodEntity(food Food) model.Food {
	return model.Food{
		Id:         food.Id,
		Name:       food.Name,
		City:       food.City,
		Country:    food.Country,
		Vegetarian: food.Vegetarian,
	}
}
