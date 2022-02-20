package http

import (
	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/validation"
)

type Rest interface {
	Start(cfg configs.ConfigImp, validator validation.Validator) error
	Shutdown() error
}
