package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/app/models"
	"github.com/ahmadabd/FoodRecommended.git/app/structs"
)

const pathUrl = "food"

func randomFoodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var food structs.Foods
		var err error
		food, err = models.GetRandomFood()
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
	switch r.Method {
	case http.MethodPost:
		var food structs.Foods
		err := json.NewDecoder(r.Body).Decode(&food)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = models.CreateFood(food); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func SetupFoodsRoutes(baseUrl string) {
	http.HandleFunc(fmt.Sprintf("/%s/%s/random", baseUrl, pathUrl), randomFoodHandler)
	http.HandleFunc(fmt.Sprintf("/%s/%s", baseUrl, pathUrl), foodHandler)
}
