package entity

import (
	"github.com/google/uuid"
)

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
