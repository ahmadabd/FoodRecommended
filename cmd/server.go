package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs/yaml"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/logger/zap"
	"github.com/ahmadabd/FoodRecommended.git/internal/pkg/validation/playground"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/gorm"

	// "github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
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
	cfg, err := yaml.GetConfig("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	f, err := makeLogFile()

	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	// dbConn, err := mysql.SetupDatabase(cfg, logger)
	dbConn, err := gorm.SetupDatabase(cfg, logger)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}

	fmt.Println("Application started.")

	foodServ := food.New(dbConn)

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

func makeLogFile() (*os.File, error) {

	if os.Getenv("CONTAINER") == "1" {
		// find path
		exPath, err := os.Executable()
		exPath = filepath.Dir(exPath)

		if err != nil {
			return nil, err
		}

		// if file dosent exist
		if _, err := os.Stat(filepath.Join(exPath, "logs", "app.log")); err != nil {
			// make logs directory
			os.Mkdir(filepath.Join(exPath, "logs"), 0755)
			os.Create(filepath.Join(exPath, "logs", "app.log"))
		}

		// make app.log file in /logs
		f, err := os.OpenFile(filepath.Join(exPath, "logs", "app.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		return f, nil
	} else {
		log.Println("CONTAINER env variable not defined.", os.Getenv("CONTAINER"))

		// make app.log file in /logs
		f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		return f, nil
	}
}
