package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (map[int][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	equations := make(map[int][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid test value: %s", parts[0])
		}

		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			numbers[i], err = strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", numStr)
			}
		}

		equations[testValue] = numbers
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return equations, nil
}

func evaluatePart1(numbers []int, index int, currentValue int, target int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	if evaluatePart1(numbers, index+1, currentValue+numbers[index], target) {
		return true
	}

	if evaluatePart1(numbers, index+1, currentValue*numbers[index], target) {
		return true
	}

	return false
}

func evaluatePart2(numbers []int, index int, currentValue int, target int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	if evaluatePart2(numbers, index+1, currentValue+numbers[index], target) {
		return true
	}

	if evaluatePart2(numbers, index+1, currentValue*numbers[index], target) {
		return true
	}

	concatenatedValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, numbers[index]))
	if evaluatePart2(numbers, index+1, concatenatedValue, target) {
		return true
	}

	return false
}

func main() {
	filename := ".\\input7.txt"

	equations, err := parseInput(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalPart1 := 0
	totalPart2 := 0

	for testValue, numbers := range equations {
		if evaluatePart1(numbers, 1, numbers[0], testValue) {
			totalPart1 += testValue
		}

		if evaluatePart2(numbers, 1, numbers[0], testValue) {
			totalPart2 += testValue
		}
	}

	fmt.Println("Part 1 Total Calibration Result:", totalPart1)
	fmt.Println("Part 2 Total Calibration Result:", totalPart2)
}
