package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/echo"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *configs.Config

func init() {
	cfg = configs.GetConfig()
	dbConn, err := mysql.SetupDatabase(cfg)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	foodServ := food.New(dbConn)
	echo.New(foodServ).Start()
}

func main() {
	fmt.Println("Application started.")

	log.Fatal(http.ListenAndServe(serverConfigs(cfg), nil))
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
