package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	Key   string `json:"Key" validate:"required"`
	Value string `json:"Value" validate:"required"`
}

type Data struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required,min=3,max=20" `
	Description Template           `json:"description,omitempty"`
}

func (t Template) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Template) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}
