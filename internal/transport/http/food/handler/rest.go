package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/validation"
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

func (r *rest) Start(cfg configs.ConfigImp, validator validation.Validator) error {
	r.echo.Validator = validator
	r.echo.Use(middleware.Recover())
	r.routing("api")

	return r.echo.Start(serverConfigs(cfg))
}

func (r *rest) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.echo.Shutdown(ctx)
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf(":%v", cfg.GetServerPort())
}
