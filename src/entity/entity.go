package entity

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID string `json:"id" mapstructure:"id"`
}

func (e *Entity) GetID() string {
	return e.ID
}

func (e *Entity) GenerateUUID() {
	e.ID = uuid.New().String()
}

type JSONEntity map[string]interface{}

func (j JSONEntity) GetID() string {
	id, ok := j["id"].(string)
	if !ok {
		return ""
	}

	return id
}

func (j JSONEntity) GenerateUUID() {
	j["id"] = uuid.New().String()
}
