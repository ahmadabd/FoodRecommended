package echo

import (
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http"
)

type rest struct {
	handler *handler
}

func New(foodSrv service.Food) http.Rest {
	return &rest{
		handler: &handler{
			food: foodSrv,
		}}
}

func (r *rest) Start() {
	r.routing("api")
}
