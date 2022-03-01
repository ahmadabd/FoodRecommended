package food_test

import (
	"context"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mock"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
)

func TestRandomFood(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	service := food.New(mock.New())

	rFood, err := service.RandomFood(ctx)

	if err != nil {
		t.Error("Expected nil, got error")
	}

	t.Log("Food Id: ", rFood.Id)
}
