package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"math/rand"
	"time"
)

// Function to generate a random channel ID
func generateRandomChannelID(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	channelID := make([]byte, length)
	for i := range channelID {
		channelID[i] = charset[rand.Intn(len(charset))]
	}
	return string(channelID)
}

// Function to continuously receive messages from the server
func receiveMessages(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected from server.")
			return
		}
		fmt.Print("Message: " + message)
	}
}

// Main function to connect to the server and start chatting
func main() {
	// Prompt the user to enter the channel ID or start a new one
	fmt.Print("Enter channel ID (or press Enter to create a new one): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channelID := strings.TrimSpace(scanner.Text())

	// Generate a random channel ID if the user input is empty
	if channelID == "" {
		channelID = generateRandomChannelID(6) // Generate a random ID of 6 characters
	}

	// Connect to the server
	server := "10.15.5.137:5050" // replace with server address if different
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send the channel ID to the server
	_, err = fmt.Fprintln(conn, channelID)
	if err != nil {
		fmt.Println("Error sending channel ID:", err)
		return
	}

	// Print the channel ID after sending it
	fmt.Println("Connected to channel:", channelID)

	// Start a goroutine to receive messages
	go receiveMessages(conn)

	// Loop to read user input and send messages to the server
	for {
		fmt.Print("You: ")
		if scanner.Scan() {
			message := scanner.Text()
			if message == "exit" {
				fmt.Println("Exiting...")
				return
			}
			// Send the message to the server
			_, err := fmt.Fprintln(conn, message)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		} else {
			fmt.Println("Error reading input.")
			return
		}
	}
}
