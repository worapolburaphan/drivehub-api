package entity

import "github.com/google/uuid"

type Car struct {
	ID       string `json:"id" mapstructure:"id"`
	Name     string `json:"name" mapstructure:"name"`
	Price    int    `json:"price" mapstructure:"price"`
	Discount int    `json:"discount" mapstructure:"discount"`
}

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GenerateUUID() {
	c.ID = uuid.New().String()
}

func NewCar(id string) *Car {
	return &Car{
		ID: id,
	}
}
