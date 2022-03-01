package food

import (
	"context"
	"strconv"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/labstack/echo/v4"
)

type fd struct {
	db repository.DB
}

func New(db repository.DB) service.Food {
	return &fd{db}
}

func (food *fd) RandomFood(ctx context.Context) (model.Food, error) {
	return food.db.GetRandomFood(ctx)
}

func (food *fd) GetFoods(ctx context.Context, c echo.Context) ([]model.Food, error) {
	return food.db.GetFoods(ctx, pagination(c))
}

func (food *fd) StoreFood(ctx context.Context, newFood model.Food) error {
	return food.db.CreateFood(ctx, newFood)
}

func pagination(c echo.Context) int {
	paginationLimit, perr := strconv.ParseInt(c.QueryParam("paginationLimit"), 10, 64)

	if perr != nil || paginationLimit == 0 {
		paginationLimit = 10
	}

	return int(paginationLimit)
}
