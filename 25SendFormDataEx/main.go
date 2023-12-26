package main

import (
	"io"
	"net/http"
	"net/url"
)

func main() {
	performPostFormReq()
}

func performPostFormReq() {
	const myUrl = "http://localhost:8080/"

	//formData
	data := url.Values{}
	data.Add("name", "Anu")
	data.Add("age", "21")
	data.Add("email", "anu@go.dev")

	resp, err:=http.PostForm(myUrl,data)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	content,err:=io.ReadAll(resp.Body)
	if err!=nil{
		panic("Error reading response body")
	}
	println(string(content)) //prints the server's response to our request
	
}