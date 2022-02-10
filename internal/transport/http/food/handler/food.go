package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
)

type handler struct {
	food service.Food
}

func (handler *handler) randomFoodHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:

		foodRes, err := handler.food.RandomFood(ctx)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foodRes)
	}
}

func (handler *handler) foodHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodPost:
		var food model.Food
		err := json.NewDecoder(r.Body).Decode(&food)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = handler.food.StoreFood(ctx, food); err != nil {
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
		var foods []model.Food
		var err error
		foods, err = handler.food.GetFoods(ctx)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foods)
	}
}
