package request

type (
	Food struct {
		Name       string `json:"name"`
		City       string `json:"city"`
		Country    string `json:"country"`
		Vegetarian uint   `json:"vegetarian"`
	}
)