package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, error := os.Open("messages.txt")
	if error != nil {
		fmt.Println("Error:", error)
	}

	// Close file
	defer file.Close()

	for {
		// Allocate memory for byte buffer.
		data := make([]byte, 8)
		// Read data into byte buffer
		n, err := file.Read(data)
		if err != nil {
			log.Fatal("error", "error", err)
			break
		}
		fmt.Printf("read: %s\n", string(data[:n]))

	}
}
