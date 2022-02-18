package handler

import (
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
	"github.com/labstack/echo/v4"
)

func (handler *handler) getRandomFoodHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		food, err := handler.food.RandomFood(c.Request().Context())

		if err != nil {
			handler.logger.Error(err.Error())

			c.JSON(http.StatusInternalServerError, response.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, castFoodToResponse(food))
	}
}
