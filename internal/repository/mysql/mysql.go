package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
)

var DbConn *sql.DB

func SetupDatabase(cfg *configs.Config) {

	var err error
	DbConn, err = sql.Open("mysql", databaseConfig(cfg))
	if err != nil {
		panic(err.Error())
	}
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