package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/myHttp"
)

type rest struct {
	handler *handler
	logger  logger.Logger
}

func New(foodSrv service.Food, logger logger.Logger) myHttp.Rest {
	return &rest{
		handler: &handler{
			food:   foodSrv,
			logger: logger,
		},
		logger: logger,
	}
}

func (r *rest) Start(cfg configs.ConfigImp) error {
	r.routing("api")

	return http.ListenAndServe(serverConfigs(cfg), nil)
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
