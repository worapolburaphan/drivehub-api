package http

import (
	"github.com/drivehub-api/src/domain/car/usecase"
	"github.com/drivehub-api/src/entity"
	"github.com/drivehub-api/src/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type handler struct {
	usecase usecase.CarUseCase
}

func NewCarHandler(ctx *gin.RouterGroup, usecase usecase.CarUseCase) {
	h := &handler{
		usecase: usecase,
	}

	ctx.GET("/cars", h.GetCars)
	ctx.POST("/cars", h.CreateCar)
	ctx.PUT("/cars/:id", h.UpdateCar)
	ctx.DELETE("/cars/:id", h.DeleteCar)
}

func (h *handler) GetCars(c *gin.Context) {
	cars, err := h.usecase.GetCars()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(
		200, utils.NewResponse(cars),
	)
}

func (h *handler) CreateCar(c *gin.Context) {
	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdCar, err := h.usecase.CreateCar(&car)
	if err != nil {
		c.JSONP(500, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(201, utils.NewResponse(createdCar))
}

func (h *handler) UpdateCar(c *gin.Context) {
	var id = c.Param("id")
	if id == "" {
		c.JSONP(400, gin.H{"error": "id is required"})
		return
	}

	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSONP(400, gin.H{"error": err.Error()})
		return
	}

	car.ID = id
	updatedCar, err := h.usecase.UpdateCar(&car)
	if err != nil {
		c.JSONP(500, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(200, utils.NewResponse(updatedCar))
}

func (h *handler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	log.Printf("id: %s", id)

	if id == "" {
		c.JSONP(400, gin.H{"error": "id is required"})
		return
	}

	err := h.usecase.DeleteCar(id)
	if err != nil {
		c.JSONP(500, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(200, nil)
}
