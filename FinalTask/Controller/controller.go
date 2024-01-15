package controllers

import (
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redisDB "TemplateUserDetailsTask/Database/Redis"
	model "TemplateUserDetailsTask/Model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type BaseHandler struct {
    MongoDB *mongodb.MongoDB
	Redis *redisDB.MyRedis
	inMemory *inmemory.InMemoryDB
	AppState *model.AppState
}

func NewBaseHandler(MongoDB *mongodb.MongoDB,Redis *redisDB.MyRedis,inMemory *inmemory.InMemoryDB,appState *model.AppState) *BaseHandler {
    return &BaseHandler{
        MongoDB: MongoDB,
		Redis : Redis,
		inMemory: inMemory,
		AppState: appState,
    }
}


func (h *BaseHandler)Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.inMemory.CreateTemplate(data)
	h.MongoDB.CreateTemplate(data)
	h.Redis.CreateTemplate(data)
	json.NewEncoder(w).Encode(data)	

}
func(h *BaseHandler)Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	vars := mux.Vars(r)
    stringData := vars["stringdata"]
	newData1 := vars["newData1"]
	newData2 := vars["newData2"]

    data := model.Data{
        Name: stringData,
        Description: model.Template{
			Key: newData1,
			Value: newData2,
		},
    }
	h.inMemory.UpdateTemplate(data)
	h.MongoDB.UpdateTemplate(data)
	h.Redis.UpdateTemplate(data)
	io.WriteString(w,"Updated!")
    json.NewEncoder(w).Encode(data)
			
}

func (h *BaseHandler)Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	vars := mux.Vars(r)
	data := vars["data"]
	h.inMemory.DeleteTemplate(data)
	h.MongoDB.DeleteTemplate(data)
	h.Redis.DeleteTemplate(data)
	json.NewEncoder(w).Encode("data")
}

func (h *BaseHandler)Refresh(w http.ResponseWriter, r *http.Request){
	h.inMemory.RefreshData(h.AppState)
	h.MongoDB.RefreshData(h.AppState)
	h.Redis.RefreshData(h.AppState)	
	fmt.Fprintf(w,"Done Refreshing Data!")
}

func (h *BaseHandler)Test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	test1,err := h.MongoDB.TestData()
	if err!=nil{
		fmt.Println(err)
	}
	
	test2,err:=h.Redis.TestData()
	if err!=nil{
		fmt.Println(err)
	}
	test3,err:=h.inMemory.TestData()
	if err!=nil{
		fmt.Println(err)
	}
	result := map[string]interface{}{"Mongo": test1,"Redis":test2,"In-Memory":test3}
	json.NewEncoder(w).Encode(result)
}

