package main

import (
	"fmt"
	"net"
)

func main(){
	listen,err:=net.Listen("tcp",":8080")
	errorHandle(err)
	defer listen.Close()
	for{
		conn,err:=listen.Accept()
		errorHandle(err)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	fmt.Println("Client Connected")
	conn.Write([]byte("Hello World \n"))
	var message string
	fmt.Println("Enter the message:")
	fmt.Scanln(&message)
	conn.Write([]byte(message))
	conn.Close()
}

func errorHandle(err error){
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
}