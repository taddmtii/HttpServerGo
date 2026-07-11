package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
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
			log.Println("Error occured when accepting connection:", error)
			continue
		}
		// After we acceot connection, make sure that if we do not hear back for 5 seconds give up.
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

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
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	// Split the request line into method, route, and HTTP version.
	requestLineParts := strings.Fields(requestLine)
	if len(requestLineParts) != 3 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		return
	}
	method, route, _ := strings.TrimSpace(requestLineParts[0]), strings.TrimSpace(requestLineParts[1]), strings.TrimSpace(requestLineParts[2])

	// Next, read the headers line by line and put them in a map.
	headerMap := make(map[string]string)

	for {
		header, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Error when reading header...")
			return
		}

		header = strings.TrimSpace(header)
		if header == "" {
			break
		}
		headerParts := strings.SplitN(header, ":", 2)
		if len(headerParts) != 2 {
			continue
		}
		key, value := strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1])

		// Regex to sanitize the string (alphanumeric)
		key_is_valid := regexp.MustCompile(`^[a-zA-Z0-9-]*$`).MatchString(key)
		// value_is_valid := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(value)

		_, exists := headerMap[key]
		if !exists && key_is_valid {
			headerMap[key] = value
		}
	}

	// Body parsing logic

	if method != "GET" {
		// Get content length number so we know how many bytes to allocate and read.
		content_length, err := strconv.Atoi(headerMap["Content-Length"])
		if err != nil {
			log.Println("Content-Length value not found.")
		}
		// Allocate buffer for body content.
		body_content := make([]byte, content_length)
		// Only read up to the content_lengths number since bodys do not have a common delimiter.
		n, err := io.ReadFull(reader, body_content)
		if err != nil {
			log.Println("Error reading bytes from reader into buffer.")
		}
	}

	if route == "/health" {
		handleHealth(conn, method)
	}
}

func handleHealth(conn net.Conn, method string) {
	if method == "GET" {
		response := "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"
		_, err := conn.Write([]byte(response))
		if err != nil {
			log.Println("write error: ", err)
		}
	} else {
		conn.Write([]byte("HTTP/1.1 405 Method Not Allowed\r\n\r\n"))
	}
}
