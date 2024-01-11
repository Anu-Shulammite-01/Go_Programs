package connections

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connect to MongoDB
func ConnectMongoDB() (*mongo.Client) {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        panic(err.Error())
    }
    return client
}
// Close the MongoDB connection
func CloseConnection(client *mongo.Client){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err!=nil{
		panic(err)
	} else {
		fmt.Println("Successfully closed Mongo DB Connection!")
	}
}
//redis connection
func ConnectRedis()(*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to Redis Server : %s\n", pong)
	return client
}
//close redis connection
func CloseRedisConn(c *redis.Client){
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Closed Redis Connection Successfully.")
	}
}

