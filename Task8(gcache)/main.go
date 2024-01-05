package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

func main(){
	fmt.Println("gcache example")
	//creates a cache with 20 items with expiration time of 5 seconds using LRU policy
	gc := gcache.New(20).Expiration(5 * time.Second).LRU().Build()
	
	gc.Set("name","Anu")
	val,err:=gc.Get("name")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n", val)

	//checking key present or not
	exists:=gc.Has("name")
	fmt.Println("Key exists? ",exists)

	//wait for 6 second
	<-time.After(6*time.Second) 
	size:=gc.Len(false)
	fmt.Println("Size:",size)

	val,_=gc.Get("name")      
	if val==nil {                   
		fmt.Println("Key is removed after expiration.")
	}else{
		fmt.Printf("Value %s \n",val)
	}
	exists = gc.Has("name")
	fmt.Println("Key exists? ",exists)

	// adding more elements
	gc.Set("name","Ana")
	gc.Set("city","Tpt")
	gc.Set("place","AP")
	keys:=gc.Keys(true)
	for _,key:=range keys{
		val,err := gc.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Println(key,":",val)
	}
	gc.Purge()
	size=gc.Len(true)
	fmt.Println("Size:",size)
}