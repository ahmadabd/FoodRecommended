package http

import "github.com/ahmadabd/FoodRecommended.git/internal/configs"

type Rest interface {
	Start(cfg configs.ConfigImp) error
	Shutdown() error
}
