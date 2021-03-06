package handler

import (
	"fmt"
)

const pathUrl = "food"

func (r *rest) routing(baseUrl string) {
	r.echo.GET(fmt.Sprintf("/%s/%s/random", baseUrl, pathUrl), r.handler.getRandomFoodHandler())
	r.echo.POST(fmt.Sprintf("/%s/%s", baseUrl, pathUrl), r.handler.createFoodHandler())
	r.echo.GET(fmt.Sprintf("/%s/%s/all", baseUrl, pathUrl), r.handler.getFoodsHandler())
}
