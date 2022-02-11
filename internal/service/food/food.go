package food

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
)

type fd struct {
	db repository.DB
}

func New(db repository.DB) service.Food {
	return &fd{db: db}
}

func (food *fd) RandomFood(ctx context.Context) (model.Food, error) {
	return food.db.GetRandomFood(ctx)
}

func (food *fd) GetFoods(ctx context.Context) ([]model.Food, error) {
	return food.db.GetFoods(ctx)
}

func (food *fd) StoreFood(ctx context.Context, newFood model.Food) error {
	return food.db.CreateFood(ctx, newFood)
}
