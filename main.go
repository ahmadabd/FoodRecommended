package main

import (
	"log"

	"github.com/ahmadabd/FoodRecommended.git/cmd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
