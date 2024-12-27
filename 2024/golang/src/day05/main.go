package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day05a := day05a("day05a.txt")
	fmt.Printf("Day 05 Part A: %d\n", day05a)
	day05b := day05b("day05b.txt")
	fmt.Printf("Day 05 Part B: %d\n", day05b)
}

func readFile(filename string) (map[int][]int, [][]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	edges := make(map[int][]int, 0)

	pastEdges := false

	input := make([][]int, 0)

	// Scan each line in file
	for scanner.Scan() {
		value := (scanner.Text())
		if value == "" {
			pastEdges = true
		} else {
			if !pastEdges {
				ints := readInts(value, "|")

				edges[ints[1]] = append(edges[ints[1]], ints[0])

			} else {
				input = append(input, readInts(value, ","))
			}
		}
	}

	return edges, input
}

func readInts(line string, seperator string) []int {
	ints := make([]int, 0)

	for _, value := range strings.Split(line, seperator) {
		int, _ := strconv.Atoi(value)
		ints = append(ints, int)
	}

	return ints
}

func day05a(filename string) int {
	edges, allPageNumbers := readFile(filename)

	result := 0

	for _, pageNumbers := range allPageNumbers {
		result += getMiddleNumber(pageNumbers, edges)
	}

	return result
}

func getMiddleNumber(pageNumbers []int, edges map[int][]int) int {
	forbiddenNumbers := make([]int, 0)

	for _, pageNumber := range pageNumbers {
		if slices.Contains(forbiddenNumbers, pageNumber) {
			return 0
		}

		numbers := edges[pageNumber]

		if numbers != nil {
			forbiddenNumbers = append(forbiddenNumbers, numbers...)
		}
	}

	return pageNumbers[(len(pageNumbers) / 2)]
}

func day05b(filename string) int {
	edges, allPageNumbers := readFile(filename)

	result := 0

	for _, pageNumbers := range allPageNumbers {
		result += getMiddleNumbersFailing(pageNumbers, edges)
	}

	return result
}

func getMiddleNumbersFailing(pageNumbers []int, edges map[int][]int) int {
	forbiddenNumbers := make([]int, 0)

	for _, pageNumber := range pageNumbers {
		if slices.Contains(forbiddenNumbers, pageNumber) {

			sorted := TopologicalSort(pageNumbers, edges)
			return sorted[(len(sorted) / 2)]
		}

		numbers := edges[pageNumber]

		if numbers != nil {
			forbiddenNumbers = append(forbiddenNumbers, numbers...)
		}

	}

	return 0
}

func TopologicalSort(input []int, edges map[int][]int) []int {
	result := make([]int, len(input))
	copy(result, input)

	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]

		// Break early if A dosen't have any edges
		if _, hasA := edges[a]; !hasA {
			return false
		}

		// Break early if B dosen't have any edges
		if _, hasB := edges[b]; !hasB {
			return true
		}

		return slices.Contains(edges[a], b)
	})

	return result
}
