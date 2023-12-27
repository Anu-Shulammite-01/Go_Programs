package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8081")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello, World!")
}
