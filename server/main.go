package main

import (
	"fmt"
	"net"
)

func handlerClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected client", conn.RemoteAddr().String())

	// Loop for read message of client
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error read ", err)
			return
		}
		clientMessage := string(buffer[:n])
		fmt.Printf("Client Message: %s\n", clientMessage)

		// Response client
		response := "Received: " + clientMessage
		conn.Write([]byte(response))
	}
}

func main() {

	fmt.Println("Server Runner...")

	// Create TCP Server on PORT 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error for initialize server", err)
		return
	}

	defer listener.Close()

	fmt.Printf("Server listenet on localhost:8080")

	// Waiting client connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}
		go handlerClient(conn) // Handle multiple clients concurrently
	}
}
