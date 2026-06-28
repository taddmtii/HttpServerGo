package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Returns a listener object that sets up a sever to listen for incoming connections on local network on port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server is now listening on port 8080...")

	for {
		// Accepts incoming connection request from queue and returns a new socket
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting a connection: %v", err)
			continue
		}

		// spins up a goroutine. Basically a lightweight thread to simplify concurrency.
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Network address (IP addr and port number from client)
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s\n", clientAddr)

	// Create a buffered reader
	reader := bufio.NewReader(conn)

	for {
		// Read until newline
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", clientAddr)
			break
		}

		fmt.Printf("Received from %s: %s", clientAddr, message)

		// Echo back to client (cast string to byte slice)
		_, err = conn.Write([]byte("Echo: " + message))
		if err != nil {
			fmt.Printf("Error writing to the cleint: %v\n", err)
			break
		}
	}
}
