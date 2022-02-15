package service

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

type Food interface {
	RandomFood(ctx context.Context) (model.Food, error)
	GetFoods(ctx context.Context, paginationLimit int) ([]model.Food, error)
	StoreFood(ctx context.Context, newFood model.Food) error
}
