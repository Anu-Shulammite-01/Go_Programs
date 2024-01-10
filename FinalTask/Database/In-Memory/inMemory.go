package inmemory

import model "TemplateUserDetailsTask/Model"

// In-memory
type InMemoryDB struct {
	data map[string]model.Template
}

func (db *InMemoryDB) CreateTemplate(key string, value model.Template) {
	db.data[key] = value
}

func (db *InMemoryDB) UpdateTemplate(oldKey string, newKey string,value model.Template) {
	_, exists := db.data[oldKey]
	if exists {
		delete(db.data, oldKey)
		db.data[newKey] = value
	}
}

func (db *InMemoryDB) DeleteTemplate(key string) {
	delete(db.data, key)
}

func (db *InMemoryDB) Refresh() error {
	db.data = make(map[string]model.Template)
	return nil
}

func (db *InMemoryDB) Test(string)([]string,error) {
	//print all keys and values in map
	keys := []string{}
	for k := range db.data{
		keys = append(keys,k)
		}
		return keys,nil
}