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
	d := 0
	h := 0
	// Initialize the hash maps
	//directions := make(map[string]int)
	//forwards := make(map[string]int)

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

		//d : = 0
		// Update the hash maps based on the command
		switch command {
		case "forward":
			h += value
		case "up":
			d -= value
		case "down":
			d += value
		default:
			log.Fatalf("Unknown command: %s", command)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := d * h
	fmt.Println(result)
}
