package food

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRandomFood(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		food, err := getRandomFood()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(food))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func handleFood(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var data Food
		json.NewDecoder(r.Body).Decode(&data)
		if err := storeNewFood(data); err != nil {
			log.Panicln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}

var foodRoute = "food"

func SetupRoute(baseRoute string) {
	http.HandleFunc(fmt.Sprintf("%s/%s/random", baseRoute, foodRoute), handleRandomFood)
	http.HandleFunc(fmt.Sprintf("%s/%s", baseRoute, foodRoute), handleFood)
}
