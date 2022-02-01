package food

import (
	"context"
	"math/rand"
	"time"

	"github.com/ahmadabd/FoodRecommended.git/database"
)

var foods []Food

func getRandomFood() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := database.DbConn.QueryContext(ctx, `SELECT 
	id,name FROM iran`)

	if err != nil {
		return "", err
	}
	defer result.Close()

	for result.Next() {
		var food Food
		err := result.Scan(&food.Id, &food.Name)
		if err != nil {
			return "", err
		}
		foods = append(foods, food)
	}

	return foods[rand.Intn(len(foods))].Name, nil
}

func storeNewFood(data Food) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := database.DbConn.ExecContext(ctx, `INSERT INTO iran(name) VALUES(?)`, data.Name)
	if err != nil {
		return err
	}
	return nil
}
