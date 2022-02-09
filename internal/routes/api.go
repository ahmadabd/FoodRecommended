package routes

import (
	"github.com/ahmadabd/FoodRecommended.git/app/controllers"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
)

const baseUrl = "api"

func SetpuRoutes(dbConn repository.DB) {
	controllers.SetupFoodsRoutes(baseUrl, dbConn)
}
