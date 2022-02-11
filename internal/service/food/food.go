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
		result = castFoodToResponse(randFood)
	}

	return result, err
}

func (food *fd) GetFoods(ctx context.Context) ([]response.Food, error) {

	var result []response.Food

	allFoods, err := food.db.GetFoods(ctx)

	if err == nil {
		for _, f := range allFoods {
			tmp := castFoodToResponse(f)
			result = append(result, tmp)
		}
	}

	return result, err
}

func (food *fd) StoreFood(ctx context.Context, newFd request.Food) error {

	newFood := castRequestToFood(newFd)

	return food.db.CreateFood(ctx, newFood)
}

func castFoodToResponse(food model.Food) response.Food {
	var tmp response.Food
	tmp.Id = food.Id
	tmp.Name = food.Name
	tmp.City = food.City
	tmp.Country = food.Country
	if food.Vegetarian == enum.Vegetarian {
		tmp.Vegetarian = "vegetarian"
	} else if food.Vegetarian == enum.NonVegetarian {
		tmp.Vegetarian = "non-vegetarian"
	}

	return tmp
}

func castRequestToFood(food request.Food) model.Food {
	var newFood model.Food

	newFood.Name = food.Name
	newFood.City = food.City
	newFood.Country = food.Country
	if food.Vegetarian == "vegetarian" {
		newFood.Vegetarian = enum.Vegetarian
	} else if food.Vegetarian == "non-vegetarian" {
		newFood.Vegetarian = enum.NonVegetarian
	}

	return newFood
}
