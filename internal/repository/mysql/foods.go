package mysql

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

func GetRandomFood(ctx context.Context) (model.Food, error) {
	var food model.Food

	result := DbConn.QueryRowContext(ctx, `SELECT 
		id, 
		name, 
		city, 
		country,
		vegetarian
		FROM foods ORDER BY RAND() LIMIT 1`)

	if err := result.Scan(&food.Id, &food.Name, &food.City, &food.Country, &food.Vegetarian); err != nil {
		return food, err
	}

	return food, nil
}

func CreateFood(ctx context.Context, food model.Food) error {
	_, err := DbConn.ExecContext(ctx,
		`INSERT INTO foods (name, city, country, vegetarian) VALUES (?, ?, ?, ?)`,
		food.Name,
		food.City,
		food.Country,
		food.Vegetarian,
	)

	return err
}

func GetFoods(ctx context.Context) ([]model.Food, error) {
	var foods []model.Food

	rows, err := DbConn.QueryContext(ctx, `SELECT 
		id, 
		name, 
		city, 
		country,
		vegetarian
		FROM foods`)
	if err != nil {
		return foods, err
	}

	defer rows.Close()

	for rows.Next() {
		var food model.Food
		if err := rows.Scan(&food.Id, &food.Name, &food.City, &food.Country, &food.Vegetarian); err != nil {
			return foods, err
		}
		foods = append(foods, food)
	}

	return foods, nil
}
