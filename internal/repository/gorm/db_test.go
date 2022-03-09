package gorm_test

import (
	"context"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomFood(t *testing.T) {
	db, cleanup := setup(t)

	defer cleanup()

	food := model.Food{
		Name:       faker.FirstName(),
		City:       faker.FirstName(),
		Country:    faker.FirstName(),
		Vegetarian: 0,
	}

	t.Run("Create New Food", func(t *testing.T) {
		err := db.CreateFood(context.TODO(), food)
		assert.Nil(t, err)
	})

	t.Run("Get Random Food", func(t *testing.T) {
		_, err := db.GetRandomFood(context.TODO())
		assert.Nil(t, err)
	})
}

func TestGetFoods(t *testing.T) {
	db, cleanup := setup(t)

	defer cleanup()

	food := model.Food{
		Name:       faker.FirstName(),
		City:       faker.FirstName(),
		Country:    faker.FirstName(),
		Vegetarian: 0,
	}

	t.Run("Create New Food", func(t *testing.T) {
		err := db.CreateFood(context.TODO(), food)
		assert.Nil(t, err)
	})

	t.Run("Get Random Food", func(t *testing.T) {
		_, err := db.GetFoods(context.TODO(), 1)
		assert.Nil(t, err)
	})
}