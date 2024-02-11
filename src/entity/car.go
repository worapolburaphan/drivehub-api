package entity

type Car struct {
	Entity
	ID       string `json:"id" mapstructure:"id"`
	Name     string `json:"name" mapstructure:"name"`
	Price    int    `json:"price" mapstructure:"price"`
	Discount int    `json:"discount" mapstructure:"discount"`
}

func NewCar(id string) *Car {
	return &Car{
		Entity: Entity{ID: id},
	}
}
