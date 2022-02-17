package handler

import (
	"net/http"
	"strconv"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/enum"
	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/service"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/request"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	food   service.Food
	logger logger.Logger
}

func (handler *handler) randomFoodHandler(c echo.Context) error {

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

func (handler *handler) addFoodHandler(c echo.Context) error {

	newFood := new(request.Food)

	if err := c.Bind(newFood); err != nil {
		handler.logger.Error(err.Error())

		c.JSON(http.StatusBadRequest, response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
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

func (handler *handler) allFoodsHandler(c echo.Context) error {

	paginationLimit, perr := strconv.ParseInt(c.QueryParam("paginationLimit"), 10, 64)
	if perr != nil || paginationLimit == 0 {
		paginationLimit = 10
	}

	var foods []model.Food
	var err error
	foods, err = handler.food.GetFoods(c.Request().Context(), int(paginationLimit))
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

func castFoodToResponse(food model.Food) *response.Food {
	var tmp response.Food
	tmp.Id = food.Id
	tmp.Name = food.Name
	tmp.City = food.City
	tmp.Country = food.Country

	if food.Vegetarian == enum.Vegetarian {
		tmp.Vegetarian = "vegetarian"
	}
	tmp.Vegetarian = "non-vegetarian"

	return &tmp
}

func castRequestToFood(food request.Food) *model.Food {
	var newFood model.Food

	newFood.Name = food.Name
	newFood.City = food.City
	newFood.Country = food.Country
	if food.Vegetarian == "vegetarian" {
		newFood.Vegetarian = enum.Vegetarian
	} else if food.Vegetarian == "non-vegetarian" {
		newFood.Vegetarian = enum.NonVegetarian
	}

	return &newFood
}
