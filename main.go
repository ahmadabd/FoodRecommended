package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/database"
	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/routes"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *configs.Config

func init() {
	cfg = configs.GetConfig()
	database.SetupDatabase(cfg)
}

func main() {
	fmt.Println("Application starting ...")

	routes.SetpuRoutes()

	log.Fatal(http.ListenAndServe(serverConfigs(cfg), nil))
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
