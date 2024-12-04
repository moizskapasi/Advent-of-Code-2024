package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([][]int, error) {
	var reports [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := []int{}
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		clonedReport := make([]int, len(report))
		copy(clonedReport, report)
		tempReport := append(clonedReport[:i], clonedReport[i+1:]...)

		if isSafe(tempReport) {
			return true
		}
	}

	return false
}

func countSafeReports(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}
	return safeCount
}

func main() {
	inputFile := "C:\\Users\\moizs\\Downloads\\AoC\\Day 2\\input2.txt"

	reports, err := parseInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	safeCount := countSafeReports(reports)
	fmt.Printf("The number of safe reports is: %d\n", safeCount)
}
