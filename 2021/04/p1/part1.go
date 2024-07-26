package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers [5][5]int
	marked  [5][5]bool
}

func (b *Board) mark(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] == number {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *Board) hasWon() bool {
	// Check rows
	for i := 0; i < 5; i++ {
		win := true
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check columns
	for j := 0; j < 5; j++ {
		win := true
		for i := 0; i < 5; i++ {
			if !b.marked[i][j] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	return false
}

func (b *Board) score(lastNumber int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum * lastNumber
}

func parseInput(scanner *bufio.Scanner) ([]int, []Board) {
	// Read the first line of numbers to draw
	scanner.Scan()
	numStrs := strings.Split(scanner.Text(), ",")
	numbers := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		numbers[i], _ = strconv.Atoi(numStr)
	}

	var boards []Board
	var currentBoard Board
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if rowIndex == 5 {
				boards = append(boards, currentBoard)
				currentBoard = Board{}
				rowIndex = 0
			}
			continue
		}

		nums := strings.Fields(line)
		for j, num := range nums {
			currentBoard.numbers[rowIndex][j], _ = strconv.Atoi(num)
		}
		rowIndex++

		// Handle the last board (EOF case)
		if rowIndex == 5 {
			boards = append(boards, currentBoard)
			currentBoard = Board{}
			rowIndex = 0
		}
	}
	return numbers, boards
}

func main() {
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers, boards := parseInput(scanner)

	for _, number := range numbers {
		for i := range boards {
			boards[i].mark(number)
			if boards[i].hasWon() {
				fmt.Printf("Board %d wins with score %d\n", i, boards[i].score(number))
				return
			}
		}
	}
}
