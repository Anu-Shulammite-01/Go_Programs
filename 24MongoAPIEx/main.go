package main

import (
	router "MongoAPIExample/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server getting started...")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4001",r))
	fmt.Println("Listening to port 4001...")
}