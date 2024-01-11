package inmemory

import model "TemplateUserDetailsTask/Model"

// In-memory
type InMemoryDB struct {
	User []model.Data
}

func (db *InMemoryDB) CreateInMemory()([]model.Data){
	db.User = make([]model.Data, 0)
	return db.User
}



func (db *InMemoryDB) CreateTemplate(data model.Data) {
	for _, value := range db.User {
		if value.Name == data.Name {
			panic("User Already Exist") 
		}
	}
	db.User = append(db.User, data)
}


func (db *InMemoryDB) UpdateTemplate(oldData model.Data,newData model.Data) {
	for index,value := range db.User{
		if value.Name == oldData.Name {
			db.User[index] = newData
			return
		}
	}
	panic("No User Found to be Updated")
}

func (db *InMemoryDB) DeleteTemplate(data model.Data) {
	for i, v := range db.User {
		if v.Name == data.Name {
			db.User = append(db.User[:i], db.User[i+1:]...)
			return
		}
	}
	panic("No User Found to be Deleted")
}

func (db *InMemoryDB) RefreshData() error {
	db.User=make([]model.Data,0)
	return nil
}

func (db *InMemoryDB) TestData()([]string,error) {
	//print all keys and values in map
	keys:= make([]string,len(db.User))
	var key string
	for i := range db.User {	key=db.User[i].Name
		keys[i]=key
	}
	return keys,nil
}