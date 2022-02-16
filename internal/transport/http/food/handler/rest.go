package handler

import (
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type rest struct {
	echo    *echo.Echo
	handler *handler
}

func New(foodSrv service.Food, logger logger.Logger) http.Rest {
	return &rest{
		echo: echo.New(),
		handler: &handler{
			food:   foodSrv,
			logger: logger,
		},
	}
}

func (r *rest) Start(cfg configs.ConfigImp) error {
	r.echo.Use(middleware.Recover())
	r.routing("api")

	return r.echo.Start(serverConfigs(cfg))
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
