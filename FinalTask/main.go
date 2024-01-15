package main

import (
	connections "TemplateUserDetailsTask/Database/Connections"
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redisDB "TemplateUserDetailsTask/Database/Redis"
	model "TemplateUserDetailsTask/Model"
	router "TemplateUserDetailsTask/Router"
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("API Example")
	//MongoDB
	mongoClient,err := connections.ConnectMongoDB()
	if err != nil{
		log.Fatal(err)
	}
	mongoDBClient:=mongodb.NewMongoDB(mongoClient)
	defer connections.CloseConnection(mongoDBClient)

	//Redis
	redisClient,err:=connections.ConnectRedis()
	if err != nil{
		log.Fatal(err)
	}
	RedisClient := redisDB.NewMyRedis(redisClient)
	defer connections.CloseRedisConn(RedisClient)

	//In-Memory
	inMemory := inmemory.NewInMemoryDB()
	
	//AppState
	appState := model.NewAppState()

	//Router
	r := router.InitializeRoutes(inMemory, mongoDBClient, redisClient, appState)
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
	}
}



