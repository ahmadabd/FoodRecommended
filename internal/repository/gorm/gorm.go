package gorm

import (
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorm struct {
	Db     *gorm.DB
	Logger logger.Logger
}

func SetupDatabase(cfg configs.ConfigImp, logger logger.Logger) (repository.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseConfig(cfg)), &gorm.Config{})
	
	if err != nil {
		logger.Error(fmt.Sprintf("Error while connecting to database: %v", err))
		return nil, err
	}
	
	return &Gorm{Db: db, Logger: logger}, err
}

func databaseConfig(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.GetDatabaseUser(),
		cfg.GetDatabasePassword(),
		cfg.GetDatabaseHost(),
		cfg.GetDatabasePort(),
		cfg.GetDatabaseName())
}
