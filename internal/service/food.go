package service

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/labstack/echo/v4"
)

type Food interface {
	RandomFood(ctx context.Context) (model.Food, error)
	GetFoods(ctx context.Context, c echo.Context) ([]model.Food, error)
	StoreFood(ctx context.Context, newFood model.Food) error
}
