package food

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/request"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
)

type fd struct {
	db repository.DB
}

func New(db repository.DB) service.Food {
	return &fd{db: db}
}

func (food *fd) RandomFood(ctx context.Context) (response.Food, error) {
	var result response.Food

	randFood, err := food.db.GetRandomFood(ctx)

	if err == nil {
		result.Id = randFood.Id
		result.Name = randFood.Name
		result.City = randFood.City
		result.Country = randFood.Country

		if randFood.Vegetarian == enum.Vegetarian {
			result.Vegetarian = "vegetarian"
		} else if randFood.Vegetarian == enum.NonVegetarian {
			result.Vegetarian = "non-vegetarian"
		}
	}

	return result, err
}

func (food *fd) GetFoods(ctx context.Context) ([]response.Food, error) {

	var result []response.Food

	allFoods, err := food.db.GetFoods(ctx)

	if err == nil {
		for _, f := range allFoods {
			var tmp response.Food
			tmp.Id = f.Id
			tmp.Name = f.Name
			tmp.City = f.City
			tmp.Country = f.Country
			if f.Vegetarian == enum.Vegetarian {
				tmp.Vegetarian = "vegetarian"
			} else if f.Vegetarian == enum.NonVegetarian {
				tmp.Vegetarian = "non-vegetarian"
			}
			result = append(result, tmp)
		}
	}

	return result, err
}

func (food *fd) StoreFood(ctx context.Context, newFd request.Food) error {
	var newFood model.Food

	newFood.Name = newFd.Name
	newFood.City = newFd.City
	newFood.Country = newFd.Country
	if newFd.Vegetarian == "vegetarian" {
		newFood.Vegetarian = enum.Vegetarian
	} else if newFd.Vegetarian == "non-vegetarian" {
		newFood.Vegetarian = enum.NonVegetarian
	}

	return food.db.CreateFood(ctx, newFood)
}
