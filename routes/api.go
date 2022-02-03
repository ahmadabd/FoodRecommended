package routes

import "github.com/ahmadabd/FoodRecommended.git/app/controllers"

const baseUrl = "api"

func SetpuRoutes() {
	controllers.SetupFoodsRoutes(baseUrl)
}
