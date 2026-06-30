package main

import (
	"bufio"
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

	// Create a reader for responses that get sent back from the server.
	// serverReader := bufio.NewReader(conn)
	// Create a reader for the client to intake request from user.
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Request: ")
		request, error := reader.ReadString('\n')
		if error != nil {
			break
		}

		// Write the request to the server.
		_, error = conn.Write([]byte(request))
		if error != nil {
			log.Fatal(error)
		}
	}
}
