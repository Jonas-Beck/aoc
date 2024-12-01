package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := day01a("day01a.txt")
	fmt.Println(result)
}

func readLists(filename string) ([]float64, []float64) {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	leftList := make([]float64, 0)
	rightList := make([]float64, 0)

	// Scan each line in file
	for scanner.Scan() {

		// init slice for parsed float value
		floats := make([]float64, 0)

		// Convert both strings to float64
		for _, field := range strings.Fields(scanner.Text()) {
			num, _ := strconv.ParseFloat(field, 64)
			floats = append(floats, num)
		}

		// append value to correct list
		leftList = append(leftList, floats[0])
		rightList = append(rightList, floats[1])
	}

	// return both lists in float64
	return leftList, rightList
}

func day01a(filename string) float64 {
	leftList, rightList := readLists(filename)

	var result float64

	// Sort both lists
	sort.Float64s(leftList)
	sort.Float64s(rightList)

	// Loop over list
	for index, value := range leftList {
		// append absolute value of smallest number - smallest number
		result += math.Abs(value - rightList[index])
	}

	return result
}

// func day01b(filename string) float64 {
// 	return 0
// }
