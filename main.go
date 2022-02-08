package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
	"github.com/ahmadabd/FoodRecommended.git/internal/routes"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *configs.Config

func init() {
	cfg = configs.GetConfig()
	mysql.SetupDatabase(cfg)
}

func main() {
	fmt.Println("Application starting ...")

	routes.SetpuRoutes()

	log.Fatal(http.ListenAndServe(serverConfigs(cfg), nil))
}

func serverConfigs(cfg configs.ConfigImp) string {
	return fmt.Sprintf("%v:%v", cfg.GetServerHost(), cfg.GetServerPort())
}
