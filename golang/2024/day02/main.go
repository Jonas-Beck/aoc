package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	day02a := day02a("day02a.txt")
	fmt.Printf("Day 02 Part A: %f\n", day02a)
	day02b := day02b("day02b.txt")
	fmt.Printf("Day 02 Part B: %f\n", day02b)
}

func readReports(filename string) [][]float64 {
	file, _ := os.Open(filename)
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// init slice for reports
	reports := make([][]float64, 0)

	// Scan each line in file
	for scanner.Scan() {

		// init slice for parsed float value
		numbers := make([]float64, 0)

		// Convert both strings to float64
		for _, field := range strings.Fields(scanner.Text()) {
			num, _ := strconv.ParseFloat(field, 64)
			numbers = append(numbers, num)
		}

		reports = append(reports, numbers)

	}

	// return both lists in float64
	return reports
}

func validDifference(previousValue float64, currentValue float64) bool {
	diff := math.Abs(previousValue - currentValue)
	return diff >= 1 && diff <= 3
}

func validDirection(previousValue float64, currentValue float64, increasingValue bool) bool {
	return (!increasingValue && currentValue < previousValue) || (increasingValue && currentValue > previousValue)
}

func day02a(filename string) float64 {
	reports := readReports(filename)

	var safeReportCount int

	for _, report := range reports {
		if isSafeReport(report) {
			safeReportCount++
		}
	}

	return float64(safeReportCount)
}

func isSafeReport(report []float64) bool {
	previousValue := report[0]
	increasingValue := report[0] < report[1]

	for _, value := range report[1:] {
		if valid := validDifference(previousValue, value); !valid {
			return false
		}

		if valid := validDirection(previousValue, value, increasingValue); !valid {
			return false
		}
		previousValue = value
	}

	return true
}

// DAY 2 PART B

func day02b(filename string) float64 {
	reports := readReports(filename)

	var safeReportCount int

	for _, report := range reports {
		if isSafeReportWithRemoval(report) {
			safeReportCount++
		}
	}

	return float64(safeReportCount)
}

func isSafeReportWithRemoval(report []float64) bool {
	if isSafeReport(report) {
		return true
	}

	// If report isn't safe
	// Loop over report length and remove level at index
	for index := range len(report) {
		if isSafeReport(RemoveIndex(report, index)) {
			return true
		}
	}

	return false
}

func RemoveIndex(s []float64, index int) []float64 {
	newSlice := make([]float64, 0, len(s)-1)
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, s[index+1:]...)
	return newSlice
}
