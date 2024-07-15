package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Initialize variables
	var prev int
	var counter int
	firstLine := true

	// Loop through the file line by line
	for scanner.Scan() {
		line := scanner.Text() // Read the current line

		// Convert the current line to an integer
		current, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err) // Handle the error if the conversion fails
		}

		// Skip comparison for the first line
		if firstLine {
			prev = current
			firstLine = false
			continue
		}

		// Compare current with previous value
		if current > prev {
			counter++
		}

		// Update prev to the current line for the next iteration
		prev = current
	}

	// Print the result
	fmt.Println(counter)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
