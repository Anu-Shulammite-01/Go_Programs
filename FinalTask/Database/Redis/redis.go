package redis

import (
	model "TemplateUserDetailsTask/Model"
	"context"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
)

// Redis
type Redis struct {
	client *redis.Client
}

func (db *Redis) CreateTemplate(user model.Data) {
	if _,err := db.client.Get(context.Background(), user.Name).Result(); err != nil {
		if err := db.client.Set(context.Background(), user.Name,user.Description, 0).Err();err != nil {
			log.Fatalf("Failed to set data to redis: %s", err)
		}
	} else {
		log.Println("User is already exists.")
	}
}

func (db *Redis) UpdateTemplate(oldData model.Data,newData model.Data) {
	db.DeleteTemplate(oldData)
	db.CreateTemplate(newData)
}

func (db *Redis) DeleteTemplate(data model.Data) {
	if res,err := db.client.Del(context.Background(), data.Name).Result();err == nil{
		if res > 0 {
			log.Printf("%d key deleted\n",res)
		}else{
			log.Println("No such key!")
		}
	}else{
		log.Printf("Failed to delete the data from redis:%s",err)
	}


}

func (db *Redis) RefreshData() error {
	return nil
}

func (db *Redis) TestData()([]string,error) {
	//get all keys and values from redis
	keys,err:= db.client.Keys(context.Background(), "*").Result()
	if err!=nil{
		return []string{},fmt.Errorf("failed to get keys from Redis :%v",err)
	}
	var result []string
	for _,key := range keys{
		value,_:=db.client.Get(context.Background(),key).Result()
		result=append(result,"Key:"+key+" Value:"+value)
	}
	return result,nil
}
