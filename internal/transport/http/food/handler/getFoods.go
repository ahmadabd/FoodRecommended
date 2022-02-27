package handler

import (
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
	"github.com/labstack/echo/v4"
)

func (handler *handler) getFoodsHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		var foods []model.Food
		var err error
		foods, err = handler.food.GetFoods(c.Request().Context(), c)
		if err != nil {
			handler.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, response.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		res := new([]response.Food)

		if err == nil {
			for _, food := range foods {
				tmp := castFoodToResponse(food)
				*res = append(*res, *tmp)
			}
		}

		return c.JSON(http.StatusOK, res)
	}
}
