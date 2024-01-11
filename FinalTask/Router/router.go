package router

import (

	controllers "TemplateUserDetailsTask/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	var control controllers.BaseHandler
	// Health Check Endpoint
	r.HandleFunc("/create", control.Create).Methods("POST")
	r.HandleFunc("/update", control.Update).Methods("PUT")
	r.HandleFunc("/delete", control.Delete).Methods("DELETE")
	r.HandleFunc("/refresh", control.Refresh).Methods("GET")
	r.HandleFunc("/test", control.Test).Methods("GET")
	return r

}