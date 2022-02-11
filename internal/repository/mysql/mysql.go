package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
)

type mysql struct {
	db *sql.DB
}

func SetupDatabase(cfg configs.ConfigImp) (repository.DB, error) {

	db, err := sql.Open("mysql", databaseConfig(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &mysql{db: db}, nil
}

func databaseConfig(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		cfg.GetDatabaseUser(),
		cfg.GetDatabasePassword(),
		cfg.GetDatabaseHost(),
		cfg.GetDatabasePort(),
		cfg.GetDatabaseName(),
	)
}
