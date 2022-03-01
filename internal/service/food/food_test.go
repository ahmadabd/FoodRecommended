package food_test

import (
	"context"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mock"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
)


func TestStoreFood(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	service := food.New(mock.New())

	food := model.Food{
		Id:         1,
		Name:       "Mock Food",
		City:       "Mock City",
		Country:    "Mock Country",
		Vegetarian: enum.Vegetarian,
	}

	err := service.StoreFood(ctx, food)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestGetFoods(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	service := food.New(mock.New())

	_, err := service.GetFoods(ctx)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestRandomFood(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	service := food.New(mock.New())

	_, err := service.RandomFood(ctx)

	if err != nil {
		t.Error("Expected nil, got error")
	}
}
