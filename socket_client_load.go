package main

import (
	"fmt"
	"net"
	"time"
	"github.com/go-faker/faker/v4"
)

func main() {
	// Connect to the server
	server := "10.15.5.137:5050" // Change this if your server is running elsewhere
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Specify a channel ID (you can modify this)
	channelID := "TESTCHANNEL"
	_, err = fmt.Fprintf(conn, "%s\n", channelID)
	if err != nil {
		fmt.Println("Error sending channel ID:", err)
		return
	}

	fmt.Println("Connected to channel:", channelID)

	// Continuously send fake messages
	for {
		// Generate a fake text message using the faker library
		message := faker.Sentence() // You can use other faker functions as well
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
