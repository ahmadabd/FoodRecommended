package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs/yaml"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/myHttp/food/handler"
	"github.com/urfave/cli/v2"
)

var ServeCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := yaml.GetConfig()
	dbConn, err := mysql.SetupDatabase(cfg)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}

	fmt.Println("Application started.")

	foodServ := food.New(dbConn)

	go func() {
		if err := handler.New(foodServ).Start(cfg); err != nil {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	return nil
}
