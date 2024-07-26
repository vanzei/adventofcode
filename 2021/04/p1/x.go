// Manually doing

package main

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
	//check rows
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
	//check columns
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
