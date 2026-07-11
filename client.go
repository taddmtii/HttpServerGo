package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Establish the connection to the server.
	conn, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		log.Fatal(error)
	}

	defer conn.Close()

	// request, err := os.ReadFile("requests/health.txt")
	request := "GET /health HTTP/1.1\r\n" +
		"Host: localhost:8080\r\n" +
		"User-Agent: Chrome\r\n" +
		"Connection: keep-alive\r\n" +
		"Content-Length: 38\r\n" +
		"\r\n" +
		"{Id: 78912, Quantity: 1, Price: 18.00}"

	// if err != nil {
	// 	fmt.Println("Error reading the request file:", err)
	// }
	fmt.Println("Sending Request... ")

	// Write the request to the server.
	_, error = conn.Write([]byte(request))
	if error != nil {
		log.Fatal(error)
	}

	// Read responses from the server.
	buffer := make([]byte, 1024)
	n, error := conn.Read(buffer)
	if error != nil {
		log.Fatal(error)
	}
	// Read the entire buffer.
	fmt.Println(string(buffer[:n]))

}
