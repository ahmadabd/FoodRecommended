package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/myHttp/food/request"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/myHttp/food/response"
)

type handler struct {
	food   service.Food
	logger logger.Logger
}

func (handler *handler) randomFoodHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:

		foodRes, err := handler.food.RandomFood(ctx)

		if err != nil {
			handler.logger.Error(fmt.Sprintf("error happen in getting random food: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var result response.Food
		result = castFoodToResponse(foodRes)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (handler *handler) foodHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodPost:
		var newFood request.Food
		err := json.NewDecoder(r.Body).Decode(&newFood)
		if err != nil {
			handler.logger.Info(fmt.Sprintf("error happen in add new food data from user: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		food := castRequestToFood(newFood)

		if err = handler.food.StoreFood(ctx, food); err != nil {
			handler.logger.Error(fmt.Sprintf("error happen in storing new food: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *handler) foodsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		paginationLimit, perr := strconv.ParseInt(r.URL.Query().Get("paginationLimit"), 10, 64)
		if perr != nil || paginationLimit == 0 {
			paginationLimit = 10
		}

		var foods []model.Food
		var err error
		foods, err = handler.food.GetFoods(ctx, int(paginationLimit))
		if err != nil {
			handler.logger.Error(fmt.Sprintf("error happen in getting all food: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var result []response.Food
		if err == nil {
			for _, food := range foods {
				tmp := castFoodToResponse(food)
				result = append(result, tmp)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func castFoodToResponse(food model.Food) response.Food {
	var tmp response.Food
	tmp.Id = food.Id
	tmp.Name = food.Name
	tmp.City = food.City
	tmp.Country = food.Country
	if food.Vegetarian == enum.Vegetarian {
		tmp.Vegetarian = "vegetarian"
	} else if food.Vegetarian == enum.NonVegetarian {
		tmp.Vegetarian = "non-vegetarian"
	}

	return tmp
}

func castRequestToFood(food request.Food) model.Food {
	var newFood model.Food

	newFood.Name = food.Name
	newFood.City = food.City
	newFood.Country = food.Country
	if food.Vegetarian == "vegetarian" {
		newFood.Vegetarian = enum.Vegetarian
	} else if food.Vegetarian == "non-vegetarian" {
		newFood.Vegetarian = enum.NonVegetarian
	}

	return newFood
}
