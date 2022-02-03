package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/database"
	"github.com/ahmadabd/FoodRecommended.git/routes"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	database.SetupDatabase()
}

func main() {
	fmt.Println("Application starting ...")

	routes.SetpuRoutes()

	log.Fatal(http.ListenAndServe(":5000", nil))
}
