package routes

import "github.com/ahmadabd/FoodRecommended.git/app/controllers"

var baseUrl = "api"

func SetpuRoutes() {
	controllers.SetupFoodsRoutes(baseUrl)
}
