package database

import (
	"database/sql"
	"fmt"

	"github.com/ahmadabd/FoodRecommended.git/configs"
)

var DbConn *sql.DB

func SetupDatabase() {

	var err error
	DbConn, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		configs.Conf.GetDatabaseUser(),
		configs.Conf.GetDatabasePassword(),
		configs.Conf.GetDatabaseHost(),
		configs.Conf.GetDatabasePort(),
		configs.Conf.GetDatabaseName()))
	if err != nil {
		panic(err.Error())
	}
}
