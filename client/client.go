package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: Client GO  <RUNNER>")
		return
	}
	message := os.Args[1]
	serverAddr := "localhost:8080"

	// Connect to server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connect to server: ", err)
		return
	}
	defer conn.Close()

	// Send message to server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Failed to send data to the server: ", err)
		return
	}

	// Read response to the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error to read response to the server: ", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Response to the server:", response)

}
