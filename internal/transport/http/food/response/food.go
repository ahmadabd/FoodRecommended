package response

type (
	Food struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		City       string `json:"city"`
		Country    string `json:"country"`
		Vegetarian uint   `json:"vegetarian"`
	}

	Error struct {
		Status int `json:"status"`
	}
)
