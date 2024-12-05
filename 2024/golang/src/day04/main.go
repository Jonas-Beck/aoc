package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	day04a := day04a("day04a.txt")
	fmt.Printf("Day 04 Part A: %f\n", day04a)
	// day04b := day04b("day04b.txt")
	// fmt.Printf("Day 04 Part B: %f\n", day04b)
}

func readFile(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var words [][]rune

	// Scan each line in file
	for scanner.Scan() {
		values := []rune(scanner.Text())
		words = append(words, values)
	}

	return words
}

// Misunderstood Day 4 A
// Task for future me to do this correct
func day04a(filename string) int {
	words := readFile(filename)

	var result int

	// Loop horizontol
	for _, row := range words {
		result += checkSlice(row)
	}

	// Loop vertical
	for colIndex := range words[0] {
		var values []rune

		for rowIndex := range words {
			values = append(values, words[rowIndex][colIndex])
		}
		result += checkSlice(values)
	}

	// Loop main diagonal
	result += loopDiagonal(words, true)
	result += loopDiagonal(words, false)

	return result
}

func loopDiagonal(words [][]rune, reverse bool) int {
	sizeRow := len(words)
	sizeCol := len(words[0])

	count := 0

	if reverse {
		// Reverse direction: Top-right to bottom-left
		for col := sizeCol - 1; col >= 0; col-- {
			var slice []rune
			for i, j := 0, col; i < sizeRow && j >= 0; i, j = i+1, j-1 {
				slice = append(slice, words[i][j])
			}
			count += checkSlice(slice)
		}
		for row := 1; row < sizeRow; row++ {
			var slice []rune
			for i, j := row, sizeCol-1; i < sizeRow && j >= 0; i, j = i+1, j-1 {
				slice = append(slice, words[i][j])
			}
			count += checkSlice(slice)
		}
	} else {
		// Normal direction: Top-left to bottom-right
		for col := 0; col < sizeCol; col++ {
			var slice []rune
			for i, j := 0, col; i < sizeRow && j < sizeCol; i, j = i+1, j+1 {
				slice = append(slice, words[i][j])
			}
			count += checkSlice(slice)
		}
		for row := 1; row < sizeRow; row++ {
			var slice []rune
			for i, j := row, 0; i < sizeRow && j < sizeCol; i, j = i+1, j+1 {
				slice = append(slice, words[i][j])
			}
			count += checkSlice(slice)
		}
	}

	return count
}

func checkSlice(words []rune) int {
	var xmasCount int

	target := []rune("XMAS")
	targetReverse := []rune("SAMX")

	var current, currentReverse int

	for _, value := range words {
		if value == target[current] {
			current++
			if current == 4 {
				xmasCount++
				current = 0
			}
		} else if value == target[0] {
			current = 1
		} else {
			current = 0
		}

		if value == targetReverse[currentReverse] {
			currentReverse++
			if currentReverse == 4 {
				xmasCount++
				currentReverse = 0
			}
		} else if value == targetReverse[0] {
			currentReverse = 1
		} else {
			currentReverse = 0
		}

	}

	return xmasCount
}

//
// func checkRune(value rune, current string, target []rune) int {
// 	if value == target[len(current)] {
// 		current += string(value)
// 		if len(current) == 4 {
// 			current = ""
// 			return 1
// 		}
// 	}
// 	return 0
// }

// func day04b(filename string) float64 {
// }
