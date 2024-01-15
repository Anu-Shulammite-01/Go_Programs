package redisDB

import (
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"

	"github.com/redis/go-redis/v9"
)

// Redis
type MyRedis struct {
	Client *redis.Client
	UpdateChan chan model.Data
	DeleteChan chan string
}

func NewMyRedis(client *MyRedis) *MyRedis {
	return &MyRedis{
		Client:      client.Client,
		UpdateChan: make(chan model.Data,100),
		DeleteChan: make(chan string,100),
	}
}

func (db *MyRedis) CreateTemplate(data model.Data)error {
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

	// Assuming model.Data has a Template field to store the processed template
	data.Description.Value = tpl.String()

	if _,err := db.Client.Get(context.Background(), data.Name).Result(); err == redis.Nil {
		bytes, err := data.Description.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal data: %w", err)
		}
		if err := db.Client.Set(context.Background(), data.Name,bytes, 0).Err();err != nil {
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

	// Assuming model.Data has a Template field to store the processed template
	data.Description.Value = tpl.String()

	ctx := context.Background()
	bytes, err := data.Description.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := db.Client.Set(ctx, data.Name, bytes, 0).Err(); err != nil {
		return fmt.Errorf("failed to update data in Redis: %w", err)
	}
	fmt.Println("Successfully updated the user details.")
	db.UpdateChan <- data
	return nil
}

func (db *MyRedis) DeleteTemplate(data string)error {
	if res,err := db.Client.Del(context.Background(), data).Result();err == nil{
		if res > 0 {
			fmt.Printf("%d key deleted\n",res)
			db.DeleteChan <- data
		}else{
			return errors.New("no such key found")
		}
	}else{
		return err
	}
	return nil
}

func (db *MyRedis) RefreshData(appState *model.AppState) {
	go func() {
		for {
			select {
			case data1 := <-db.UpdateChan:
				appState.Templates[data1.Name] = data1.Description
				fmt.Printf("Updated appState; Key: %s, Template: %+v\n", data1.Name, data1.Description)
			case data2 := <-db.DeleteChan:
				delete(appState.Templates, data2)
				fmt.Printf("Deleted from appState; Key: %s\n", data2)
			}
		}
	}()
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
