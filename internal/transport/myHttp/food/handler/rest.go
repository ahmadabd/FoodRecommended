package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/myHttp"
)

type rest struct {
	handler *handler
}

func New(foodSrv service.Food) myHttp.Rest {
	return &rest{
		handler: &handler{
			food: foodSrv,
		}}
}

func (r *rest) Start(cfg configs.ConfigImp) error {
	r.routing("api")

	return http.ListenAndServe(serverConfigs(cfg), nil)
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
