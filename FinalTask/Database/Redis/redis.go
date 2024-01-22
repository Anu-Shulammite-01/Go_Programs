package redisDB

import (
	mongodb "TemplateUserDetailsTask/Database/MongoDB"
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

type MyRedis struct {
	Client *redis.Client
}

func NewMyRedis(client *MyRedis) *MyRedis {
	return &MyRedis{
		Client:      client.Client,
	}
}

func (db *MyRedis) CreateTemplate(data model.Data)error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	data.Description.Value = tpl.String()

	if _,err := db.Client.Get(context.Background(), data.Name).Result(); err == redis.Nil {
		bytes, err := data.Description.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal data: %w", err)
		}
		if err := db.Client.Set(context.Background(), data.Name,bytes, 10 *time.Second).Err();err != nil {
			return fmt.Errorf("failed to set data to redis: %w", err)
		}
		fmt.Println("Sucessfully created template in Redis!")
	}else if err != nil {
		return fmt.Errorf("failed to get data from redis: %w", err)
	} else {
		return fmt.Errorf("user already exists in Redis")
	}
	return nil
}

func (db *MyRedis) UpdateTemplate(data model.Data)error {
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	data.Description.Value = tpl.String()

	ctx := context.Background()
	bytes, err := data.Description.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := db.Client.Set(ctx, data.Name, bytes, 10 *time.Second).Err(); err != nil {
		return fmt.Errorf("failed to update data in Redis: %w", err)
	}
	fmt.Println("Successfully updated the user details.")
	return nil
}

func (db *MyRedis) DeleteTemplate(data string)error {
	if res,err := db.Client.Del(context.Background(), data).Result();err == nil{
		if res > 0 {
			fmt.Printf("%d key deleted\n",res)
		}else{
			return errors.New("no such key found")
		}
	}else{
		return err
	}
	return nil
}

func (db *MyRedis) RefreshData(mongoClient *mongodb.MongoDB, data string) error {
	//context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongoClient.Client.Database("UserInfo").Collection("Details")
	filter := bson.D{{Key: "name", Value: data}}

	// Retrieve the document from MongoDB
	var result model.Data
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return fmt.Errorf("error getting data from MongoDB: %s", err)
	}

	// Write the value to Redis
	err = db.Client.Set(ctx, result.Name, result.Description, 0).Err()
	if err != nil {
		fmt.Printf("Error writing data to Redis: %s\n", err)
		return err
	}

	fmt.Printf("Refreshed Redis database; Key: %s, Value: %s\n", result.Name, result.Description.Key+":"+result.Description.Value)

	return nil
}

func (db *MyRedis) TestData() ([]string, error) {
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
