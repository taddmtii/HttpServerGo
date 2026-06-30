package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Start listening for connections on specific port.
	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Server is now listening on port 8080")

	// Make sure that the connection gets closed as some point.
	defer listener.Close()

	for {
		// Accept each connection from the queue as it comes through.
		conn, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
			break
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make sure the connection gets closed.
	defer conn.Close()

	// This is the HTTP request split into a []byte, or stream of bytes.
	reader := bufio.NewReader(conn)

	// Read the request line of the request.
	// Example: GET /plans HTTP/1.1 \r\n
	requestLine, _ := reader.ReadString('\n')

	parts := strings.Split(requestLine, " ")

	method, route, version := parts[0], parts[1], parts[2]

	fmt.Printf("Method: %s ---- Route: %s ---- Version %s", method, route, version)

}
