package connections

import (
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	redisDB "TemplateUserDetailsTask/Database/Redis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connect to MongoDB
func ConnectMongoDB() (*mongodb.MongoDB,error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Anu_Shulammite:Password@cluster0.kwjddrn.mongodb.net/")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %s", err)
    }
	fmt.Println("Successfully connected to MongoDB")
	return &mongodb.MongoDB{Client:client},nil
}

// Close the MongoDB connection
func CloseConnection(c *mongodb.MongoDB)error{
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := c.Client.Disconnect(ctx); err!=nil{
		return fmt.Errorf("failed to close MongoDB connection: %s", err)
	} else {
		fmt.Println("Successfully closed Mongo DB Connection!")
	}
	return nil
}

//redis connection
func ConnectRedis()(*redisDB.MyRedis,error){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %s", err)
	}
	fmt.Printf("Connected to Redis Server : %s\n", pong)
	return &redisDB.MyRedis{Client:client},nil
}

//close redis connection
func CloseRedisConn(c *redisDB.MyRedis)error{
	if err := c.Client.Close(); err != nil {
		return fmt.Errorf("failed to close Redis connection: %s", err)
	}
	fmt.Println("Closed Redis Connection Successfully.")
	return nil
}
