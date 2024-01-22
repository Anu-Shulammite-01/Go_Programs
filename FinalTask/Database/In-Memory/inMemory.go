package inmemory

import (
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type InMemoryDB struct {
	User map[string]model.Template
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		User: make(map[string]model.Template),
	}
}

func (db *InMemoryDB) CreateTemplate(data model.Data)error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	
	data.Description.Value = tpl.String()

	for key := range db.User {
		if key == data.Name {
			return fmt.Errorf("user already exists")
		}
	}
	db.User[data.Name] =  data.Description
	fmt.Println("Successfully created template in In-Memory!")
	return nil
}


func (db *InMemoryDB) UpdateTemplate(data model.Data)error {
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	data.Description.Value = tpl.String()

	_, ok := db.User[data.Name]
	if !ok {
		return fmt.Errorf("template does not exist")
	}
	db.User[data.Name] = data.Description
	fmt.Printf("Successfully updated the details of %s.\n", data.Name)
	return nil
}

func (db *InMemoryDB) DeleteTemplate(data string)error {
	_, ok := db.User[data]
	if !ok {
		return fmt.Errorf("no such user found")
	}else{
		delete(db.User, data)
		fmt.Printf("Successfully deleted %v.\n", data)

	}
	return nil
}



func (db *InMemoryDB) TestData()([]string,error) {
	// for keys, values := range db.User{
	// 	fmt.Println(keys," : ",values)
	// }
	var results []string
	for key,value := range db.User {
		results = append(results, key+" : "+value.Key+" = "+value.Value)
	}
	return results ,nil
}