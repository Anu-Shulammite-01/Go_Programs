package controllers

import (
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redis "TemplateUserDetailsTask/Database/Redis"
	model "TemplateUserDetailsTask/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseHandler struct {
    MongoDB *mongodb.MongoDB
	Redis *redis.Redis
	inMemory *inmemory.InMemoryDB
}

func NewBaseHandler(MongoDB *mongodb.MongoDB,Redis *redis.Redis,inMemory *inmemory.InMemoryDB) *BaseHandler {
    return &BaseHandler{
        MongoDB: MongoDB,
		Redis : Redis,
		inMemory: inMemory,
    }
}

func (h *BaseHandler)Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	h.inMemory.CreateTemplate(data)
	h.MongoDB.CreateTemplate(data)
	h.Redis.CreateTemplate(data)
	json.NewEncoder(w).Encode(data)	

}
func(h *BaseHandler)Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	var oldData model.Data
	var newData model.Data
	err := json.NewDecoder(r.Body).Decode(&oldData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the UpdateTemplate function
	h.MongoDB.UpdateTemplate(oldData, newData)
	h.Redis.UpdateTemplate(oldData, newData)
	h.inMemory.UpdateTemplate(oldData,newData)
	json.NewEncoder(w).Encode(newData)
}

func (h *BaseHandler)Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	
	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.MongoDB.DeleteTemplate(data)
	h.Redis.DeleteTemplate(data)
	h.inMemory.DeleteTemplate(data)
	json.NewEncoder(w).Encode("data")
}

func (h *BaseHandler)Refresh(w http.ResponseWriter, r *http.Request){
	h.inMemory.RefreshData()
	fmt.Fprintf(w,"Done Refreshing Data!")
}

func (h *BaseHandler)Test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	test1,err := h.MongoDB.TestData()
	if err!=nil{
		fmt.Println(err)
	}
	test2,_:=h.Redis.TestData()
	test3,_:=h.inMemory.TestData()
	result := map[string]interface{}{"mongo": test1,"redis":test2,"In-Memory":test3}
	json.NewEncoder(w).Encode(result)
}

