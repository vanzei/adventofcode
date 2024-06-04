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
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Read all lines into a slice
	var lines []int
	for scanner.Scan() {
		line := scanner.Text() // Read the current line
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err) // Handle the error if the conversion fails
		}
		lines = append(lines, value)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Initialize the counter
	var counter int

	// Loop through the slice with a window of size 3
	for i := 0; i < len(lines)-3; i++ {
		// Calculate the sums of the current window and the next window
		sum1 := lines[i] + lines[i+1] + lines[i+2]
		sum2 := lines[i+1] + lines[i+2] + lines[i+3]

		// Compare the sums
		if sum2 > sum1 {
			counter++
		}
	}

	// Print the result
	fmt.Println(counter)
}
