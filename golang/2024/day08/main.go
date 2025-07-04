package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jonas-Beck/aoc/golang/lib/grid"
)

func main() {
	day08a := day08a("day08a.txt")
	fmt.Printf("Day 07 Part A: %d\n", day08a)
	day08b := day08b("day08b.txt")
	fmt.Printf("Day 06 Part B: %d\n", day08b)
}

func readFile(filename string) (grid.Grid, map[rune][]grid.Position) {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	types := make(map[rune][]grid.Position)

	row := 0
	col := 0

	// Scan each line in file
	for scanner.Scan() {
		line := scanner.Text()
		col = len(line)

		for pos, char := range line {
			if char != '.' {
				types[char] = append(types[char], grid.Position{Row: row, Col: pos})
			}
			row++
		}
	}

	return grid.NewGrid(row, col, '.'), types
}

// Calculate positions with Vector Subtractions and multiply with 2 to get the 2:1 ratio
func day08a(filename string) int {
	board, types := readFile(filename)

	uniqueAntiNodes := make(map[grid.Position]bool)

	for _, positions := range types {
		uniqueCombinationCallback(positions, func(x grid.Position, y grid.Position) {
			// (2x₁ - x₂, 2y₁ - y₂)
			antinode1 := grid.Position{
				Row: 2*x.Row - y.Row,
				Col: 2*x.Col - y.Col,
			}

			if valid := !board.CheckOutOfBounds(antinode1); valid {
				uniqueAntiNodes[antinode1] = true
			}

			// (2x₂ - x₁, 2y₂ - y₁)
			antinode2 := grid.Position{
				Row: 2*y.Row - x.Row,
				Col: 2*y.Col - x.Col,
			}

			if valid := !board.CheckOutOfBounds(antinode2); valid {
				uniqueAntiNodes[antinode2] = true
			}
		})
	}

	return len(uniqueAntiNodes)
}

func day08b(filename string) int {
	board, types := readFile(filename)

	uniqueAntiNodes := make(map[grid.Position]bool)

	for _, positions := range types {
		uniqueCombinationCallback(positions, func(x grid.Position, y grid.Position) {
			displacement := grid.Position{
				Row: y.Row - x.Row, // from x to y
				Col: y.Col - x.Col,
			}

			uniqueAntiNodes[x] = true
			uniqueAntiNodes[y] = true

			for i := 1; ; i++ {

				antinode1 := grid.Position{
					Row: x.Row - i*displacement.Row,
					Col: x.Col - i*displacement.Col,
				}

				if board.CheckOutOfBounds(antinode1) {
					break
				}

				uniqueAntiNodes[antinode1] = true

			}

			for i := 1; ; i++ {

				antinode2 := grid.Position{
					Row: y.Row + i*displacement.Row,
					Col: y.Col + i*displacement.Col,
				}

				if board.CheckOutOfBounds(antinode2) {
					break
				}

				uniqueAntiNodes[antinode2] = true

			}
		})
	}

	return len(uniqueAntiNodes)
}

func uniqueCombinationCallback(positions []grid.Position, callback func(x grid.Position, y grid.Position)) {
	for i := 0; i < len(positions)-1; i++ {
		for j := i + 1; j < len(positions); j++ {
			x := positions[i]
			y := positions[j]
			callback(x, y)
		}
	}
}
