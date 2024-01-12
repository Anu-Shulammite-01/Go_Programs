package model

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Data struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" `
	Description Template           `json:"description,omitempty"`
}

type AppState struct {
    Templates map[string]Template
}

func NewAppState() *AppState {
    return &AppState{
        Templates: make(map[string]Template),
    }
}

func (t Template) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Template) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}
