package inmemory

import (
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"fmt"
	"text/template"
)

type InMemoryDB struct {
	User map[string]model.Template
	UpdateChan chan model.Data
	DeleteChan chan string
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		User: make(map[string]model.Template),
		UpdateChan: make(chan model.Data,100),
		DeleteChan: make(chan string,100),
	}
}

func (db *InMemoryDB) CreateTemplate(data model.Data)error {
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

	for _, value := range db.User {
		if value.Key == data.Name {
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
	db.UpdateChan <- data
	return nil
}

func (db *InMemoryDB) DeleteTemplate(data string)error {
	_, ok := db.User[data]
	if !ok {
		return fmt.Errorf("no such user found")
	}else{
		delete(db.User, data)
		fmt.Printf("Successfully deleted %v.\n", data)
		db.DeleteChan <- data
	}
	return nil
}

func (db *InMemoryDB) RefreshData(appState *model.AppState) {
	go func() {
		for {
			select {
			case data1 := <-db.UpdateChan:
				appState.Templates[data1.Name] = data1.Description
				fmt.Printf("Updated appState; Key: %s, Template: %+v\n", data1.Name, data1.Description)
			case data2 := <-db.DeleteChan:
				delete(appState.Templates, data2)
				fmt.Printf("Deleted from appState; Key: %s\n", data2)
			}
		}
	}()
}

func (db *InMemoryDB) TestData()([]string,error) {
	for keys, values := range db.User{
		fmt.Println(keys," : ",values)
	}
	var results []string
	for key,value := range db.User {
		results = append(results, key+" : "+value.Key+" = "+value.Value)
	}
	return results ,nil
}