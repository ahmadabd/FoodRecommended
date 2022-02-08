package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
)

const pathUrl = "food"

func randomFoodHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		var food model.Food
		var err error
		food, err = mysql.GetRandomFood(ctx)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(food)
	}
}

func foodHandler(w http.ResponseWriter, r *http.Request) {
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
		if err = mysql.CreateFood(ctx, food); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func foodsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		var foods []model.Food
		var err error
		foods, err = mysql.GetFoods(ctx)
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

func SetupFoodsRoutes(baseUrl string) {
	http.HandleFunc(fmt.Sprintf("/%s/%s/random", baseUrl, pathUrl), randomFoodHandler)
	http.HandleFunc(fmt.Sprintf("/%s/%s", baseUrl, pathUrl), foodHandler)
	http.HandleFunc(fmt.Sprintf("/%s/%s/all", baseUrl, pathUrl), foodsHandler)
}
