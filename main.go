package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadabd/FoodRecommended.git/database"
	"github.com/ahmadabd/FoodRecommended.git/food"
	_ "github.com/go-sql-driver/mysql"
)

var baseRoute = "/api"

func main() {
	fmt.Println("Program is starting ...")

	database.SetupDatabase()
	food.SetupRoute(baseRoute)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
