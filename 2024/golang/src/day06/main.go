package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Position struct {
	Row int
	Col int
}

func (p Position) CheckOutOfBounds(guardMap [][]rune) bool {
	// Check row bounds
	if p.Row < 0 || p.Row >= len(guardMap) {
		return true
	}

	// Check column bounds
	if p.Col < 0 || p.Col >= len(guardMap[p.Row]) {
		return true
	}

	return false
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) GetOffset() (int, int) {
	switch d {
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	default:
		return 0, 0
	}
}

func (d Direction) Next() Direction {
	return Direction((int(d) + 1) % 4)
}

func main() {
	day06a := day06a("day06a.txt")
	fmt.Printf("Day 06 Part A: %d\n", day06a)
	// day06b := day06b("day06b.txt")
	// fmt.Printf("Day 06 Part B: %d\n", day06b)
}

func readFile(filename string) ([][]rune, Position) {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var guardMap [][]rune
	var startPosition Position

	// Scan each line in file
	for scanner.Scan() {
		values := []rune(scanner.Text())
		fmt.Println(string(values))
		if index := slices.Index(values, '^'); index != -1 {
			startPosition = Position{len(guardMap), index}
		}
		guardMap = append(guardMap, values)
	}

	return guardMap, startPosition
}

func day06a(filename string) int {
	guardMap, startPosition := readFile(filename)

	visitedPositions := make(map[Position]bool)
	moveGuard(guardMap, UP, startPosition, visitedPositions)

	return len(visitedPositions)
}

func moveGuard(guardMap [][]rune, direction Direction, currentPosition Position, visitedPositions map[Position]bool) {
	visitedPositions[currentPosition] = true
	offsetRow, offsetCol := direction.GetOffset()

	newPosition := Position{Col: currentPosition.Col + offsetCol, Row: currentPosition.Row + offsetRow}

	if outOfBounds := newPosition.CheckOutOfBounds(guardMap); outOfBounds {
		return
	}

	if guardMap[newPosition.Row][newPosition.Col] == '#' {
		direction = direction.Next()
		newPosition = currentPosition
	}

	moveGuard(guardMap, direction, newPosition, visitedPositions)
}

// func day06b(filename string) int {
// 	edges, allPageNumbers := readFile(filename)
//
// 	result := 0
//
// 	for _, pageNumbers := range allPageNumbers {
// 		result += getMiddleNumbersFailing(pageNumbers, edges)
// 	}
//
// 	return result
// }
//
