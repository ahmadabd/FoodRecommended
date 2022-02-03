package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/app/models"
	"github.com/ahmadabd/FoodRecommended.git/app/structs"
)

var pathUrl = "food"

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

func SetupFoodsRoutes(baseUrl string) {
	http.HandleFunc(fmt.Sprintf("/%s/%s/random", baseUrl, pathUrl), randomFoodHandler)
}
