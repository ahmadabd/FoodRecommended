package handler

import (
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/request"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
	"github.com/labstack/echo/v4"
)

func (handler *handler) createFoodHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		newFood := new(request.Food)

		if err := c.Bind(newFood); err != nil {
			handler.logger.Error(err.Error())

			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, response.Error{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		if err := c.Validate(newFood); err != nil {
			return err
		}

		food := castRequestToFood(*newFood)

		if err := handler.food.StoreFood(c.Request().Context(), *food); err != nil {
			handler.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, response.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, response.Error{
			Status:  http.StatusCreated,
			Message: "Success",
		})
	}
}
