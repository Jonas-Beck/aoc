package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/Jonas-Beck/aoc/golang/lib/direction"
	"github.com/Jonas-Beck/aoc/golang/lib/grid"
)

type State struct {
	Pos grid.Position
	Dir direction.Direction
}

func main() {
	day06a := day06a("day06a.txt")
	fmt.Printf("Day 06 Part A: %d\n", day06a)
	day06b := day06b("day06b.txt")
	fmt.Printf("Day 06 Part B: %d\n", day06b)
}

func readFile(filename string) (grid.Grid, grid.Position) {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var guardMap grid.Grid
	var startPosition grid.Position

	// Scan each line in file
	for scanner.Scan() {
		values := []rune(scanner.Text())
		if index := slices.Index(values, '^'); index != -1 {
			startPosition = grid.Position{Row: len(guardMap), Col: index}
		}
		guardMap = append(guardMap, values)
	}

	return guardMap, startPosition
}

func day06a(filename string) int {
	guardMap, startPosition := readFile(filename)

	visitedPositions := moveGuard(guardMap, direction.UP, startPosition)

	return len(visitedPositions)
}

func moveGuard(guardMap grid.Grid, direction direction.Direction, currentPosition grid.Position) map[grid.Position]bool {
	visitedPositions := make(map[grid.Position]bool)

	for {
		visitedPositions[currentPosition] = true
		offsetRow, offsetCol := direction.GetOffset()

		newPosition := grid.Position{Col: currentPosition.Col + offsetCol, Row: currentPosition.Row + offsetRow}

		if outOfBounds := guardMap.CheckOutOfBounds(newPosition); outOfBounds {
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

func checkIfGuardStuck(guardMap grid.Grid, direction direction.Direction, currentPosition grid.Position) bool {
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

		newPosition := grid.Position{Col: currentPosition.Col + offsetCol, Row: currentPosition.Row + offsetRow}

		if outOfBounds := guardMap.CheckOutOfBounds(newPosition); outOfBounds {
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

	visitedPositions := moveGuard(guardMap, direction.UP, startPosition)

	// Remove starting position
	delete(visitedPositions, startPosition)

	obstructionPositons := 0

	for position := range visitedPositions {

		newGuardMap := guardMap.DeepCopy()
		newGuardMap[position.Row][position.Col] = '#'

		if stuck := checkIfGuardStuck(newGuardMap, direction.UP, startPosition); stuck {
			obstructionPositons++
		}

	}

	return obstructionPositons
}
