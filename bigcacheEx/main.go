package main

import (
	"context"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

func main(){
	//creating bigcache with lifewindow of 10 minutes
	cache,err:= bigcache.New(context.Background(),bigcache.DefaultConfig(10*time.Minute))
	if err!=nil{
		panic(err)
	}
	defer cache.Close()
	//setting value in the cache
	key := "Name"
	value := "Anu"
	cache.Set(key,[]byte(value))
	cache.Set("City",[]byte("Tpt"))
	//getting the value from the cache
	entry,_:=cache.Get(key)
	fmt.Println("Value for key",key,"is ",string(entry))
	entry1,_:=cache.Get("City")
	fmt.Println("Value for key City is ", string(entry1))

	//checking for nonexisting key
	entry2,err:=cache.Get("Age")
	if err==bigcache.ErrEntryNotFound {
		fmt.Println("Key Age does not exist.")
	}else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Value for key Age is %s\n", string(entry2))
	}

	err = cache.Delete("City")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = cache.Get("City")
	if err == bigcache.ErrEntryNotFound {
		fmt.Println("Successfully deleted entry for key City")
	}

	// Reset (clear) the cache
	cache.Reset()

	fmt.Println("Cache operations completed successfully")
}