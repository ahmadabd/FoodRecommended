package handler

import (
	"fmt"
	"net/http"
)

const pathUrl = "food"

func (r *rest) routing(baseUrl string) {
	http.HandleFunc(fmt.Sprintf("/%s/%s/random", baseUrl, pathUrl), r.handler.randomFoodHandler)
	http.HandleFunc(fmt.Sprintf("/%s/%s", baseUrl, pathUrl), r.handler.foodHandler)
	http.HandleFunc(fmt.Sprintf("/%s/%s/all", baseUrl, pathUrl), r.handler.foodsHandler)
}
