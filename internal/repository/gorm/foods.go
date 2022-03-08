package gorm

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

func (db *Gorm) GetRandomFood(ctx context.Context) (model.Food, error) {
	var foodEntity Food

	if err := db.Db.WithContext(ctx).Order("RAND()").First(&foodEntity).Error; err != nil {
		return model.Food{}, err
	}

	return mapToFoodEntity(foodEntity), nil
}

func (db *Gorm) CreateFood(ctx context.Context, food model.Food) error {
	f := mapFromFoodEntity(food)

	if err := db.Db.WithContext(ctx).Create(&f).Error; err != nil {
		return err
	}

	return nil
}

func (db *Gorm) GetFoods(ctx context.Context, paginationLimit int) ([]model.Food, error) {

	var foodsEntitry []Food

	if err := db.Db.WithContext(ctx).Limit(paginationLimit).Find(&foodsEntitry).Error; err != nil {
		return nil, err
	}

	var foods []model.Food

	for _, food := range foodsEntitry {
		foods = append(foods, mapToFoodEntity(food))
	}

	return foods, nil
}