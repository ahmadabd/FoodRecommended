package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs/yaml"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger/zap"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/validation/playground"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
	"github.com/ahmadabd/FoodRecommended.git/internal/service/food"
	"github.com/ahmadabd/FoodRecommended.git/internal/transport/http/food/handler"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var ServeCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := yaml.GetConfig()

	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	dbConn, err := mysql.SetupDatabase(cfg, logger)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}

	fmt.Println("Application started.")

	foodServ := food.New(dbConn, logger)

	validator := playground.New()

	go func() {
		if err := handler.New(foodServ, logger).Start(cfg, validator); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	return nil
}
