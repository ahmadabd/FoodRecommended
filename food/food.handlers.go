package food

import (
	"fmt"
	"net/http"
)

func handleRandomFood(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		food, err := GetRandomFood()
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

var foodRoute = "food"

func SetupRoute(baseRoute string) {
	http.HandleFunc(fmt.Sprintf("%s/%s/random", baseRoute, foodRoute), handleRandomFood)
}
