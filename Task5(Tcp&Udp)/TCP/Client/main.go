package main

import (
	// "bufio"
	"fmt"
	"io"
	"net"
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

	fmt.Println("Received from server:\n", string(message))

	


	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	message := scanner.Text()
	// 	fmt.Println("Received from server:\n", message)
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }
	
}