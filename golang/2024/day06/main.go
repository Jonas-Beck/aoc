package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/Jonas-Beck/aoc/golang/lib/direction"
)

type Position struct {
	Row int
	Col int
}

func CreateDeepCopy(original [][]rune) [][]rune {
	newSlice := make([][]rune, len(original))
	for i := range original {
		newSlice[i] = make([]rune, len(original[i]))
		copy(newSlice[i], original[i])
	}
	return newSlice
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

type State struct {
	Pos Position
	Dir direction.Direction
}

func main() {
	day06a := day06a("day06a.txt")
	fmt.Printf("Day 06 Part A: %d\n", day06a)
	day06b := day06b("day06b.txt")
	fmt.Printf("Day 06 Part B: %d\n", day06b)
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
		if index := slices.Index(values, '^'); index != -1 {
			startPosition = Position{len(guardMap), index}
		}
		guardMap = append(guardMap, values)
	}

	return guardMap, startPosition
}

func day06a(filename string) int {
	guardMap, startPosition := readFile(filename)

	visitedPositions := moveGuard(guardMap, UP, startPosition)

	return len(visitedPositions)
}

func moveGuard(guardMap [][]rune, direction Direction, currentPosition Position) map[Position]bool {
	visitedPositions := make(map[Position]bool)

	for {
		visitedPositions[currentPosition] = true
		offsetRow, offsetCol := direction.GetOffset()

		newPosition := Position{Col: currentPosition.Col + offsetCol, Row: currentPosition.Row + offsetRow}

		if outOfBounds := newPosition.CheckOutOfBounds(guardMap); outOfBounds {
			break
		}

		if guardMap[newPosition.Row][newPosition.Col] == '#' {
			direction = direction.Next()
		} else {
			currentPosition = newPosition
		}

	}

	return visitedPositions
}

func checkIfGuardStuck(guardMap [][]rune, direction Direction, currentPosition Position) bool {
	visitedObstructions := make(map[State]bool)

	for {
		offsetRow, offsetCol := direction.GetOffset()

		state := State{
			Pos: currentPosition,
			Dir: direction,
		}

		if visitedObstructions[state] {
			return true
		}

		newPosition := Position{Col: currentPosition.Col + offsetCol, Row: currentPosition.Row + offsetRow}

		if outOfBounds := newPosition.CheckOutOfBounds(guardMap); outOfBounds {
			return false
		}

		if guardMap[newPosition.Row][newPosition.Col] == '#' {
			visitedObstructions[state] = true
			direction = direction.Next()
		} else {
			currentPosition = newPosition
		}

	}
}

func day06b(filename string) int {
	guardMap, startPosition := readFile(filename)

	visitedPositions := moveGuard(guardMap, UP, startPosition)

	// Remove starting position
	delete(visitedPositions, startPosition)

	obstructionPositons := 0

	for position := range visitedPositions {

		newGuardMap := CreateDeepCopy(guardMap)
		newGuardMap[position.Row][position.Col] = '#'

		if stuck := checkIfGuardStuck(newGuardMap, UP, startPosition); stuck {
			obstructionPositons++
		}

	}

	return obstructionPositons
}
