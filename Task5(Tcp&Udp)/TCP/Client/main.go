package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	message, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received from server: ", string(message))
	
}
