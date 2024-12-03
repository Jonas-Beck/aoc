package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	day03a := day03a("day03a.txt")
	fmt.Printf("Day 03 Part A: %f\n", day03a)
	day03b := day03b("day03b.txt")
	fmt.Printf("Day 03 Part B: %f\n", day03b)
}

func readFile(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var corruptedMemory string

	// Scan each line in file
	for scanner.Scan() {
		corruptedMemory += scanner.Text()
	}

	return corruptedMemory
}

func day03a(filename string) float64 {
	corruptedMemory := readFile(filename)
	mulPattern := `mul\((?P<first>[0-9]+),(?P<second>[0-9]+)\)`
	mulRegex := regexp.MustCompile(mulPattern)
	var result float64

	matches := mulRegex.FindAllStringSubmatch(corruptedMemory, -1)

	for _, value := range matches {
		first, _ := strconv.ParseFloat(value[1], 64)
		second, _ := strconv.ParseFloat(value[2], 64)
		result += first * second
	}

	return result
}

func day03b(filename string) float64 {
	corruptedMemory := readFile(filename)
	mulPattern := `mul\((?P<first>[0-9]+),(?P<second>[0-9]+)\)|do\(\)|don't\(\)`
	mulRegex := regexp.MustCompile(mulPattern)
	var result float64
	var disabled bool

	matches := mulRegex.FindAllStringSubmatch(corruptedMemory, -1)

	for _, value := range matches {
		if strings.Contains(value[0], "don't()") {
			disabled = true
		} else if strings.Contains(value[0], "do()") {
			disabled = false
		}

		result += multiply(value[1], value[2], disabled)

	}

	return result
}

func multiply(firstString string, secondString string, disabled bool) float64 {
	if disabled {
		return 0
	}

	first, _ := strconv.ParseFloat(firstString, 64)
	second, _ := strconv.ParseFloat(secondString, 64)

	return first * second
}
