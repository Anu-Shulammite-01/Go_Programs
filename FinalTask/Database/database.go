package database

import(
	model "FinalTask/Model"
)

type database interface {
	CreateTemplate(key string, value Template)
	UpdateTemplate(oldKey string, newKey string,value Template)
	DeleteTemplate(key string)
	Refresh() error
	Test(string)([]string,error)
}
