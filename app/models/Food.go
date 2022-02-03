package models

import (
	"context"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/app/structs"
	"github.com/ahmadabd/FoodRecommended.git/database"
)

func GetRandomFood() (structs.Foods, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var food structs.Foods

	result := database.DbConn.QueryRowContext(ctx, `SELECT 
		id, 
		name, 
		city, 
		country 
		FROM foods ORDER BY RAND() LIMIT 1`)

	if err := result.Scan(&food.Id, &food.Name, &food.City, &food.Country); err != nil {
		return food, err
	}

	return food, nil
}
