package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	routes "recipe.core/internal/transport/rest/v1"
	"recipe.core/internal/transport/rest/v1/middlewares"
)

func Run() {
	setupLogger()

	r := gin.New()

	r.Use(
		gin.Recovery(),
		middlewares.Logger(),
	)

	routes.SetupRoutes(r)

	_ = r.Run()
}

func setupLogger() {
	f, err := os.Create("vars/logs/log.log")
	if err != nil {
		log.Println(err)
		return
	}

	gin.DefaultWriter = io.Writer(f)
}
