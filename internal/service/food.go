package service

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/request"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
)

type Food interface {
	RandomFood(ctx context.Context) (response.Food, error)
	GetFoods(ctx context.Context) ([]response.Food, error)
	StoreFood(ctx context.Context, newFood request.Food) error
}
