package main

import (
	carHttp "github.com/drivehub-api/src/domain/car/delivery/http"
	"github.com/drivehub-api/src/domain/car/repository"
	"github.com/drivehub-api/src/domain/car/usecase"
	"github.com/drivehub-api/src/external"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)
import "github.com/gin-contrib/cors"

func main() {
	app := gin.Default()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	app.Use(cors.Default())
	app.Use(gin.Logger())

	apiV1 := app.Group("/api/v1")

	apiV1.GET(
		"/check-health", func(c *gin.Context) {
			c.JSON(
				200, gin.H{
					"message": "I'm alive",
				},
			)
		},
	)

	jsonStorage := external.NewJsonStorage("./db/cars.json")

	carHttp.NewCarHandler(apiV1, usecase.NewCarUsecase(repository.NewCarRepository(jsonStorage)))

	err := app.Run(":8080")
	if err != nil {
		logger.Fatalf("could not start server: %v", err)
		return
	}
}
