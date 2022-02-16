package food

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
)

type fd struct {
	db     repository.DB
	logger logger.Logger
}

func New(db repository.DB, logger logger.Logger) service.Food {
	return &fd{db, logger}
}

func (food *fd) RandomFood(ctx context.Context) (model.Food, error) {
	return food.db.GetRandomFood(ctx)
}

func (food *fd) GetFoods(ctx context.Context, paginationLimit int) ([]model.Food, error) {
	return food.db.GetFoods(ctx, paginationLimit)
}

func (food *fd) StoreFood(ctx context.Context, newFood model.Food) error {
	return food.db.CreateFood(ctx, newFood)
}
