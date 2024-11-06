package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	"github.com/go-faker/faker/v4"
)

// clientSimulator simulates a single client connection with a dynamic channel ID, sending messages at regular intervals.
func clientSimulator(server string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Generate a unique channel ID for each client
	channelID := faker.UUIDHyphenated()

	// Connect to the server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send the dynamic channel ID
	_, err = fmt.Fprintf(conn, "%s\n", channelID)
	if err != nil {
		fmt.Println("Error sending channel ID:", err)
		return
	}

	fmt.Printf("Connected to channel: %s\n", channelID)

	// Continuously send fake messages
	for {
		// Generate a fake text message
		message := faker.Sentence()
		_, err := fmt.Fprintf(conn, "%s\n", message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		fmt.Println("Sent:", message)

		// Wait for 1 second before sending the next message
		time.Sleep(1 * time.Second)
	}
}

func main() {
	server := "10.15.5.137:5050" // Change this if your server is running elsewhere
	numClients := 1000             // Force connect 100 clients at once

	var wg sync.WaitGroup
	wg.Add(numClients)

	// Start exactly 100 client goroutines with dynamic channel IDs
	for i := 0; i < numClients; i++ {
		go clientSimulator(server, &wg)
	}

	// Wait for all client goroutines to complete
	wg.Wait()
}
