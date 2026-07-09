package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Establish the connection to the server.
	conn, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		log.Fatal(error)
	}

	defer conn.Close()

	request, err := os.ReadFile("requests/health.txt")
	if err != nil {
		fmt.Println("Error reading the request file:", err)
	}
	fmt.Println("Sending Request... ")

	for {

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
}
