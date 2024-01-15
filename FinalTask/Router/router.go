package router

import (
	controllers "TemplateUserDetailsTask/Controller"
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redisDB "TemplateUserDetailsTask/Database/Redis"
	model "TemplateUserDetailsTask/Model"

	"github.com/gorilla/mux"
)


func InitializeRoutes(inMemory *inmemory.InMemoryDB, mongoClient *mongodb.MongoDB, redisClient *redisDB.MyRedis,appState *model.AppState) *mux.Router {
	// Creating a new BaseHandler with the provided inmemory, mongoClient, and redisClient
	handler := controllers.NewBaseHandler(mongoClient, redisClient, inMemory, appState)

	// Creating a new router
	r := mux.NewRouter()

	// Setting up the routes
	r.HandleFunc("/create", handler.Create).Methods("POST")
	r.HandleFunc("/update/{stringdata}/{newData1}/{newData2}/", handler.Update).Methods("PUT")
	r.HandleFunc("/delete/{data}/", handler.Delete).Methods("DELETE")
	r.HandleFunc("/refresh", handler.Refresh).Methods("GET")
	r.HandleFunc("/test", handler.Test).Methods("GET")
	return r
}
