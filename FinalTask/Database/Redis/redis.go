package redis

import "github.com/go-redis/redis"

// Redis
type Redis struct {
	client *redis.Client
}

func (db *Redis) CreateTemplate(key string, value Template) {
	// Implementation here
}

func (db *Redis) UpdateTemplate(oldKey string, newKey string,value Template) {
	// Implementation here
}

func (db *Redis) DeleteTemplate(key string) {
	// Implementation here
}

func (db *Redis) Refresh() error {
	// Implementation here
	return nil
}

func (db *Redis) Test(string)([]string,error) {
	// Implementation here
	return nil, nil
}
