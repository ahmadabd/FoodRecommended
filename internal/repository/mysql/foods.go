package mysql

import (
	"context"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
)

func (db *Mysql) GetRandomFood(ctx context.Context) (model.Food, error) {
	var food model.Food

	result := db.Db.QueryRowContext(ctx, `SELECT 
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

func (db *Mysql) CreateFood(ctx context.Context, food model.Food) error {
	_, err := db.Db.ExecContext(ctx,
		`INSERT INTO foods (name, city, country, vegetarian) VALUES (?, ?, ?, ?)`,
		food.Name,
		food.City,
		food.Country,
		food.Vegetarian,
	)

	return err
}

func (db *Mysql) GetFoods(ctx context.Context, paginationLimit int) ([]model.Food, error) {
	var foods []model.Food

	rows, err := db.Db.QueryContext(ctx, `SELECT 
		id, 
		name, 
		city, 
		country,
		vegetarian
		FROM foods LIMIT ?`, paginationLimit)
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
