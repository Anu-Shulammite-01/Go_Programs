package main

import "fmt"

type inMemoryStore struct{
	data map[string]string
}

func(i *inMemoryStore)Set(key string,value string){
	i.data[key]=value	
}

func (i *inMemoryStore)Get(key string)string{
	return i.data[key]
}

func (i *inMemoryStore)Exists(key string)bool{
	_,ok:=i.data[key]
	return ok
}

func(i *inMemoryStore)Delete(key string){
	if i.Exists(key){
		delete(i.data, key)
	}else{
		panic("Key does not exist")
	}
}

func (i *inMemoryStore)GetAll()(map[string]string){
	return i.data
}

func main(){
	fmt.Println("In-Memory example")
	store := &inMemoryStore{data: make(map[string]string)}
	store.Set("name","Anu")
	store.Set("city","Tpt")
	fmt.Println("Name : ", store.Get("name"))
	fmt.Println("Does city exists? ", store.Exists("city"))
	getAllElements := store.GetAll()
	for k , v  := range getAllElements {
		fmt.Printf("%s = %s\n",k,v)
	}
	store.Delete("city")
	fmt.Println("Does city exists? ", store.Exists("city"))
}
