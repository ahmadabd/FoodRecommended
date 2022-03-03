package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
)

type Mysql struct {
	Db     *sql.DB
	Logger logger.Logger
}

func SetupDatabase(cfg configs.ConfigImp, logger logger.Logger) (repository.DB, error) {

	db, err := sql.Open("mysql", databaseConfig(cfg))
	if err != nil {
		logger.Error(fmt.Sprintf("Error while connecting to database: %v", err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Mysql{Db: db, Logger: logger}, nil
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
