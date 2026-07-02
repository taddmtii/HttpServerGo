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

		fmt.Println("Client has connected.")

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
	requestLineParts := strings.Split(requestLine, " ")
	method, route, _ := strings.TrimSpace(requestLineParts[0]), strings.TrimSpace(requestLineParts[1]), strings.TrimSpace(requestLineParts[2])

	headerMap := make(map[string]string)

	for {
		header, _ := reader.ReadString('\n')
		if header == "\n" {
			break
		}

		// Regular Expression that matches only strings with alpha characters)
		// regex := regexp.Compile()

		headerParts := strings.Split(header, ":")
		key, value := headerParts[0], headerParts[1]
		_, exists := headerMap[key]

		if !exists {
			headerMap[key] = value
		}
	}

	for key, value := range headerMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	if route == "/health" {
		handleHealth(conn, method)
	}
}

func handleHealth(conn net.Conn, method string) {
	if method == "GET" {
		response := "HTTP/1.1 200 OK\n"
		_, err := conn.Write([]byte(response))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conn.Write([]byte("This method is not supported at this endpoint."))
	}
}
