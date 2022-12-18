package app

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"recipe.core/internal/config"
	"recipe.core/internal/storage/psql"
	routes "recipe.core/internal/transport/rest/v1"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	file, err := os.OpenFile(
		"./vars/logs/server.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error DB connection: %v", err)
	}

	defer psql.Close(appConfig.Store.GetPgxConn())

	appConfig.App.Use(logger.New(logger.Config{
		Output:     file,
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Europe/Moscow",
	}))

	//app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	routes.SetupRoutes(appConfig)

	_ = appConfig.App.Listen(":3000")
}
