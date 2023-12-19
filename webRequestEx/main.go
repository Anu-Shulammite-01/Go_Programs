package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const Url = "https://lco.dev"

func main(){
	// GET Request
	resp,err:=http.Get(Url)
	if err!=nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status:",resp.Status)
	body,err:=io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	sb:=string(body)
	fmt.Println("Content Length: ",len(body))
	fmt.Println(sb)

	//POST Request
	post_Body,err:=json.Marshal(map[string]string{
		"name":"Anu Shulammite",
		"email":"Anu@dev.com",
	})
	if err != nil {
		log.Fatalf("Error creating json payload :%v", err)
	}
	resp_body:=bytes.NewBuffer(post_Body)
	resp1,err:=http.Post("https://postman-echo.com/post","application/json",resp_body)
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s", err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Body from POST Method:")
	body1,err:=io.ReadAll(resp1.Body)
	if err != nil {
		log.Fatalf("Error reading the response body of Post method:%v", err)
	}
	fmt.Println(string(body1))
}