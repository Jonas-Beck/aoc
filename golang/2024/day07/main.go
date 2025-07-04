package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	day07a := day07a("day07a.txt")
	fmt.Printf("Day 07 Part A: %d\n", day07a)
	day07b := day07b("day07b.txt")
	fmt.Printf("Day 06 Part B: %d\n", day07b)
}

func readFile(filename string) map[int][]int {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := make(map[int][]int)

	// Scan each line in file
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")

		key, _ := strconv.Atoi(split[0])

		numbers[key] = readInts(split[1])

	}

	return numbers
}

func readInts(line string) []int {
	ints := make([]int, 0)

	for _, value := range strings.Fields(line) {
		int, _ := strconv.Atoi(value)
		ints = append(ints, int)
	}

	return ints
}

func calculateEquation(target int, numbers []int, operators []string) int {
	var ints []int

	ints = getOperatorResults(target, numbers[1], []int{numbers[0]}, operators)

	for _, number := range numbers[2:] {
		ints = getOperatorResults(target, number, ints, operators)
	}

	if present := slices.Contains(ints, target); present {
		return target
	}

	return 0
}

func getOperatorResults(target, number int, numbers []int, operators []string) []int {
	result := make([]int, 0)

	for _, value := range numbers {
		for _, operator := range operators {
			operatorResult := applyOperator(value, number, operator)

			// Filter out numbers that are to big
			if operatorResult <= target {
				result = append(result, operatorResult)
			}
		}
	}

	return result
}

func applyOperator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	case "||":
		result, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		return result
	default:
		return 0
	}
}

func day07a(filename string) int {
	numbers := readFile(filename)

	result := 0

	operators := []string{"+", "*"}

	for target, values := range numbers {
		result += calculateEquation(target, values, operators)
	}

	return result
}

func day07b(filename string) int {
	numbers := readFile(filename)

	result := 0

	operators := []string{"+", "*", "||"}

	for target, values := range numbers {
		result += calculateEquation(target, values, operators)
	}

	return result
}
