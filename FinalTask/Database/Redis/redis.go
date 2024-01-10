package redis

import (
	model "TemplateUserDetailsTask/Model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/redis/go-redis/v9"
)

// Redis
type Redis struct {
	client *redis.Client
}

func (db *Redis) CreateTemplate(key string, value model.Template) {
	err := db.client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set data to redis: %s", err)
	}
}

func (db *Redis) UpdateTemplate(oldKey string, newKey string,value model.Template) {
	db.DeleteTemplate(oldKey)
	db.CreateTemplate(newKey, value)
}

func (db *Redis) DeleteTemplate(key string) {
	err := db.client.Del(context.Background(), key).Err()
	if err != nil && err != redis.Nil{
		log.Println(err)
		return
	}
}

func (db *Redis) Refresh() error {
	return nil
}

func (db *Redis) Test(string)([]string,error) {
	//get all keys and values from redis
	keys, err := db.client.Keys(context.Background(), "*").Result()
	if err != nil {
		return []string{}, err
	}
	values:=make(map[string]model.Template)
	for _,k:=range keys{
		val:=db.GetTemplate(k)
		values[k]=val
	}
	var res []string
	for k,v:=range values{
		res=append(res,"Key:"+k+" Value:"+fmt.Sprintf("%#v", v))
	}
	sort.Strings(res)
	return res,nil
}
func (db *Redis) GetTemplate(key string)model.Template{
	t,err:=db.client.Get(context.Background(),key).Result()
	if err==redis.Nil{
		return model.Template{}
	}else if err!=nil{
		log.Fatal(err)
	}
	var tpl model.Template
	json.Unmarshal([]byte(t),&tpl)
	return tpl
}
