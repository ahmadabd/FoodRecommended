package food_test

import (
	"context"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
	"github.com/ahmadabd/FoodRecommended.git/mocks"
	"github.com/golang/mock/gomock"
)

// func TestStoreFood(t *testing.T) {

// 	service := food.New(mock.New())

// 	food := model.Food{
// 		Id:         1,
// 		Name:       "Mock Food",
// 		City:       "Mock City",
// 		Country:    "Mock Country",
// 		Vegetarian: enum.Vegetarian,
// 	}

// 	err := service.StoreFood(context.TODO(), food)

// 	if err != nil {
// 		t.Error("Expected no error, got ", err)
// 	}
// }

// func TestGetFoods(t *testing.T) {

// 	service := food.New(mock.New())

// 	_, err := service.GetFoods(context.TODO())

// 	if err != nil {
// 		t.Error("Expected no error, got ", err)
// 	}
// }

// func TestRandomFood(t *testing.T) {

// 	service := food.New(mock.New())

// 	_, err := service.RandomFood(context.TODO())

// 	if err != nil {
// 		t.Error("Expected nil, got error")
// 	}
// }

func TestStoreFood(t *testing.T) {

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockDB := mocks.NewMockDB(mockctl)

	service := food.New(mockDB)

	food := model.Food{
		Id:         1,
		Name:       "Mock Food",
		City:       "Mock City",
		Country:    "Mock Country",
		Vegetarian: enum.Vegetarian,
	}

	mockDB.EXPECT().CreateFood(gomock.Any(), food).Return(nil).Times(1)

	err := service.StoreFood(context.TODO(), food)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestGetFoods(t *testing.T) {

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockDB := mocks.NewMockDB(mockctl)

	mockDB.EXPECT().GetFoods(gomock.Any(), 10).Return([]model.Food{
		{
			Id:         1,
			Name:       "Mock Food1",
			City:       "Mock City1",
			Country:    "Mock Country1",
			Vegetarian: enum.Vegetarian,
		},
		{
			Id:         2,
			Name:       "Mock Food2",
			City:       "Mock City2",
			Country:    "Mock Country2",
			Vegetarian: enum.NonVegetarian,
		},
	}, nil).
		Times(1)

	service := food.New(mockDB)

	_, err := service.GetFoods(context.TODO(), 10)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestRandomFood(t *testing.T) {

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockDB := mocks.NewMockDB(mockctl)

	mockDB.EXPECT().GetRandomFood(gomock.Any()).Return(model.Food{
		Id:         1,
		Name:       "Mock Food",
		City:       "Mock City",
		Country:    "Mock Country",
		Vegetarian: enum.Vegetarian,
	}, nil).
		Times(1)

	service := food.New(mockDB)

	_, err := service.RandomFood(context.TODO())

	if err != nil {
		t.Error("Expected nil, got error")
	}
}
