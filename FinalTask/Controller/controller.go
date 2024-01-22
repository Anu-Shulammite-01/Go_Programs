package controllers

import (
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redisDB "TemplateUserDetailsTask/Database/Redis"
	model "TemplateUserDetailsTask/Model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type BaseHandler struct {
    MongoDB *mongodb.MongoDB
	Redis *redisDB.MyRedis
	inMemory *inmemory.InMemoryDB
}

func NewBaseHandler(MongoDB *mongodb.MongoDB,Redis *redisDB.MyRedis,inMemory *inmemory.InMemoryDB) *BaseHandler {
    return &BaseHandler{
        MongoDB: MongoDB,
		Redis : Redis,
		inMemory: inMemory,
    }
}


func (h *BaseHandler)Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	validate := validator.New()

	// Validating the struct
	err = validate.Struct(data)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
	
		errorMessage := map[string]string{"error": fmt.Sprintf("Validation error: %s", errors)}

		// Converting the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Writing the JSON error message to the response
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonError)
		return
	}

	err = h.inMemory.CreateTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Converting the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Writing the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	err = h.MongoDB.CreateTemplate(data)
	if err != nil {
		h.inMemory.DeleteTemplate(data.Name)
		errorMessage := map[string]string{"error": err.Error()}
		jsonError, _ := json.Marshal(errorMessage)

		// Writing the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	err= h.Redis.CreateTemplate(data)
	if err!=nil{
		h.inMemory.DeleteTemplate(data.Name)
		h.MongoDB.DeleteTemplate(data.Name)
		errorMessage := map[string]string{"error": err.Error()}

		// Converting the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Writing the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	json.NewEncoder(w).Encode(data)	
}

func(h *BaseHandler)Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	err = h.inMemory.UpdateTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Converting the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Writing the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}

	err = h.MongoDB.UpdateTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}

	err = h.Redis.UpdateTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
    json.NewEncoder(w).Encode(data)		
}

func (h *BaseHandler)Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	vars := mux.Vars(r)
	data := vars["data"]

	err := h.inMemory.DeleteTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}

	err = h.MongoDB.DeleteTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}

	err = h.Redis.DeleteTemplate(data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	response := map[string]string{
        "message": "Deleted Successfully!",
    }
    // Encoding the response data as JSON and write it to the response writer
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *BaseHandler)Refresh(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","GET")
	vars := mux.Vars(r)
	data := vars["data"]
	err := h.MongoDB.RefreshData(h.inMemory,data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	err = h.Redis.RefreshData(h.MongoDB,data)
	if err != nil {
		errorMessage := map[string]string{"error": err.Error()}

		// Convert the map to JSON
		jsonError, _ := json.Marshal(errorMessage)

		// Write the JSON error message to the response
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonError)
		return
	}
	response := map[string]string{
        "message": "Done Refreshing Data!",
    }
    // Encoding the response data as JSON and write it to the response writer
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func (h *BaseHandler)Test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

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