package inmemory

import (
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"fmt"
	"sync"
	"text/template"
)

// In-memory
type InMemoryDB struct {
	User map[string]model.Template
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		User: make(map[string]model.Template),
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

	// Assuming model.Data has a Template field to store the processed template
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

	// Assuming model.Data has a Template field to store the processed template
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

func (db *InMemoryDB) RefreshData(appState *model.AppState) error {
	var wg sync.WaitGroup

	// For each key-value pair in the User map, update your application's state
	for key, value := range db.User {
		wg.Add(1)
		go func(key string, value model.Template) {
			defer wg.Done()

			// Update the application's state with the new template
			appState.Templates[key] = value
			fmt.Printf("From In-Memory ; Key: %s, Template: %+v\n", key, value)
		}(key, value)
	}

	wg.Wait()
	return nil
}

func (db *InMemoryDB) TestData()([]string,error) {
	//print all keys and values in map
	for keys, values := range db.User{
		fmt.Println(keys," : ",values)
	}
	var results []string
	for key,value := range db.User {
		results = append(results, key+" : "+value.Key+" = "+value.Value)
	}
	return results ,nil
}