package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Initialize the hash maps
	directions := make(map[string]int)
	forwards := make(map[string]int)

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop through the file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Read and trim the current line
		if line == "" {
			continue // Skip empty lines
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Invalid input format: %s", line)
		}

		command := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err) // Handle the error if the conversion fails
		}

		// Update the hash maps based on the command
		switch command {
		case "forward":
			forwards["forward"] += value
		case "up":
			directions["up"] += value
		case "down":
			directions["down"] += value
		default:
			log.Fatalf("Unknown command: %s", command)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the net up-down value
	netUpDown := directions["down"] - directions["up"]

	// Get the forward value
	forwardValue := forwards["forward"]

	// Print the result of both hash maps multiplied by the end values
	result := netUpDown * forwardValue
	fmt.Println(result)
}
