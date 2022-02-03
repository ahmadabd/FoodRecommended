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

func CreateFood(food structs.Foods) error {
	ctx, cance := context.WithTimeout(context.Background(), 10*time.Second)
	defer cance()
	_, err := database.DbConn.ExecContext(ctx, `INSERT INTO foods (name, city, country) VALUES (?, ?, ?)`, food.Name, food.City, food.Country)
	return err
}
