package database

import model "TemplateUserDetailsTask/Model"

type database interface {
	CreateTemplate(key string, value model.Template)
	UpdateTemplate(oldKey string, newKey string,value model.Template)
	DeleteTemplate(key string)
	Refresh() error
	Test(string)([]string,error)
}
