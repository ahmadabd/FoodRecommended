package handler

import (
	"net/http"
	"strconv"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
	"github.com/labstack/echo/v4"
)

func (handler *handler) getFoodsHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		paginationLimit, perr := strconv.ParseInt(c.QueryParam("paginationLimit"), 10, 64)

		var foods []model.Food
		var err error

		if perr != nil {
			foods, err = handler.food.GetFoods(c.Request().Context())
		} else {
			foods, err = handler.food.GetFoods(c.Request().Context(), int(paginationLimit))
		}

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
