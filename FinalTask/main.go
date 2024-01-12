package main

import (
	connections "TemplateUserDetailsTask/Database/Connections"
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	model "TemplateUserDetailsTask/Model"
	router "TemplateUserDetailsTask/Router"
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("API Example")
	//establishing database connection of mongoDB and redis from connections package
	mongoClient,err := connections.ConnectMongoDB()
	if err != nil{
		log.Fatal(err)
	}
	redisClient,err:=connections.ConnectRedis()
	if err != nil{
		log.Fatal(err)
	}
	db := inmemory.NewInMemoryDB()
	defer connections.CloseConnection(mongoClient)
	defer connections.CloseRedisConn(redisClient)
	appState := model.NewAppState()
	r := router.InitializeRoutes(db, mongoClient, redisClient, appState)
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
	}
}



