package database

import "database/sql"

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:XXXXXXXXXXXXX@tcp(127.0.0.1:3306)/food")
	if err != nil {
		panic(err.Error())
	}
}
