package model

type Template struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Data struct{
	Name        string       `json:"Name" validate:"required"`
	Description  *Template    `json:"Description,omitempty"`
}
