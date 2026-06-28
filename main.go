package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the server (localhost:8080)")

	// Send messages from stdin
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		// Read from the user
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// Send that message over to the server
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}

		// Now we can read the response from the server.
		response, err := serverReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Server: %s\n", response)
	}
}
