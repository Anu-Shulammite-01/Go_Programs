package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)


func Update(ctx context.Context){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://lco.dev", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err:=io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
func main(){
	ch := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i:=0 ; i<3 ; i++{
		go Update(ctx)
		ch <- true
	}
	for i:=0 ; i<3 ; i++{
		<-ch
	}
	return
}