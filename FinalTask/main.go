package main

import (
	connections "TemplateUserDetailsTask/Database/Connections"
	router "TemplateUserDetailsTask/Router"
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("API Example")
	//establishing database connection of mongoDB and redis from connections package
	mongoClient := connections.ConnectMongoDB()
	redisClient:=connections.ConnectRedis()
	defer connections.CloseConnection(mongoClient)
	defer connections.CloseRedisConn(redisClient)
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8081", r))
}



