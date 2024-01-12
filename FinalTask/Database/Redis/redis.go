package redisDB

import (
	model "TemplateUserDetailsTask/Model"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// Redis
type MyRedis struct {
	Client *redis.Client
}


func (db *MyRedis) CreateTemplate(data model.Data) {
	if _,err := db.Client.Get(context.Background(), data.Name).Result(); err == redis.Nil {
		bytes, err := data.Description.MarshalBinary()
		if err != nil {
			log.Fatalf("Failed to marshal data: %s", err)
		}
		if err := db.Client.Set(context.Background(), data.Name,bytes, 0).Err();err != nil {
			log.Fatalf("Failed to set data to redis: %s", err)
		}
		fmt.Println("Sucessfully created template in Redis!")
	}else if err != nil {
		log.Fatalf("Failed to get data from redis: %s", err)
	} else {
		log.Println("User is already exists.")
	}
}

func (db *MyRedis) UpdateTemplate(data model.Data) {
	ctx := context.Background()
	bytes, err := data.Description.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal data: %s", err)
	}
	if err := db.Client.Set(ctx, data.Name, bytes, 0).Err(); err != nil {
		log.Fatalf("Failed to update data in redis: %s", err)
	}
	fmt.Println("Successfully updated the user details.")
}

func (db *MyRedis) DeleteTemplate(data string) {
	if res,err := db.Client.Del(context.Background(), data).Result();err == nil{
		if res > 0 {
			log.Printf("%d key deleted\n",res)
		}else{
			log.Println("No such key!")
		}
	}else{
		log.Printf("Failed to delete the data from redis:%s",err)
	}


}

func (db *MyRedis) RefreshData(appState *model.AppState) error {
    // Fetch all keys from Redis
	ctx := context.Background()
	keys, err := db.Client.Keys(ctx, "*").Result()
	if err != nil {
		return fmt.Errorf("failed to get keys from Redis: %v", err)
	}

	var wg sync.WaitGroup

	// For each key, fetch the associated value and update your application's state
	for _, key := range keys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			value, err := db.Client.Get(ctx, key).Result()
			if err != nil {
				fmt.Printf("failed to get value from Redis: %v\n", err)
				return
			}

			// Unmarshal the value into a Template object
			var template model.Data
			if err := template.Description.UnmarshalBinary([]byte(value)); err != nil {
				fmt.Printf("failed to unmarshal data: %s\n", err)
				return
			}

			// Update the application's state with the new template
			appState.Templates[key] = template.Description
			fmt.Printf("FromRedis; Key: %s, Template: %+v\n", key, template)
		}(key)
	}

	wg.Wait()

	return nil
}


func (db *MyRedis) TestData() ([]string, error) {
	//want both keys and values
	ctx := context.Background()
	keys, err := db.Client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys from Redis: %v", err)
	}

	var result []string
	for _, key := range keys {
		value, err := db.Client.Get(ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to get value from Redis: %v", err)
		}
		var template model.Template
		if err := template.UnmarshalBinary([]byte(value)); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data: %s", err)
		}
		result = append(result,key+" : "+template.Key+" = "+template.Value)
	}
	return result, nil
}
