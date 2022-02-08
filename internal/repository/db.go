package repository

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

type DB interface {
	CreateFood(ctx context.Context, food model.Food) error
	GetRandomFood(ctx context.Context) (model.Food, error)
	GetFoods(ctx context.Context) ([]model.Food, error)
}
