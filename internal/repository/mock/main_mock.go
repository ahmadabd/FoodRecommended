package mock

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
)

type mock struct{}

func New() repository.DB {
	return &mock{}
}

func (m *mock) CreateFood(ctx context.Context, food model.Food) error {
	return nil
}

func (m *mock) GetRandomFood(ctx context.Context) (model.Food, error) {

	food := model.Food{
		Id:         1,
		Name:       "Mock Food",
		City:       "Mock City",
		Country:    "Mock Country",
		Vegetarian: enum.Vegetarian,
	}

	return food, nil
}

func (m *mock) GetFoods(ctx context.Context, paginationLimit int) ([]model.Food, error) {
	foods := []model.Food{
		{
			Id:         1,
			Name:       "Mock Food1",
			City:       "Mock City1",
			Country:    "Mock Country1",
			Vegetarian: enum.Vegetarian,
		},
		{
			Id:         2,
			Name:       "Mock Food2",
			City:       "Mock City2",
			Country:    "Mock Country2",
			Vegetarian: enum.NonVegetarian,
		},
	}

	return foods, nil
}
