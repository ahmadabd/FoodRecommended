package request

type (
	Food struct {
		Name       string `json:"name" validate:"required"`
		City       string `json:"city" validate:"required"`
		Country    string `json:"country" validate:"required"`
		Vegetarian string `json:"vegetarian" validate:"required"`
	}
)
