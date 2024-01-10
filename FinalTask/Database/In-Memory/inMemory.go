package inmemory

import(
	
)

// In-memory
type InMemoryDB struct {
	data map[string]string
}

func (db *InMemoryDB) CreateTemplate(key string, value Template) {
	// Implementation here
}

func (db *InMemoryDB) UpdateTemplate(oldKey string, newKey string,value Template) {
	// Implementation here
}

func (db *InMemoryDB) DeleteTemplate(key string) {
	// Implementation here
}

func (db *InMemoryDB) Refresh() error {
	// Implementation here
	return nil
}

func (db *InMemoryDB) Test(string)([]string,error) {
	// Implementation here
	return nil, nil
}