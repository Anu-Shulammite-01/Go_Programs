package router

import (
	controller "MongoAPIExample/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// Health Check Endpoint
	r.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie",controller.DeleteAlMovies).Methods("DELETE")
	

	return r
}