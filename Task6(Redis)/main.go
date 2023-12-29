package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main(){
	fmt.Println("Redis Example")
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
	
	err=client.Set(ctx,"Name","Anu",0).Err()
	if err!=nil{
		panic(err)
	}
	val,err:=client.Get(ctx,"Name").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("Name = ",val)
	defer client.Close()
}