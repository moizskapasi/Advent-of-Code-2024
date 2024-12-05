package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]int, []int, error) {
	leftList := []int{}
	rightList := []int{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input line: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}

func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func calculateSimilarityScore(leftList, rightList []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range rightList {
		frequencyMap[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * frequencyMap[num]
	}

	return similarityScore
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	inputFile := ".\\input1.txt"

	leftList, rightList, err := parseInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	totalDistance := calculateTotalDistance(leftList, rightList)
	fmt.Printf("The total distance is: %d\n", totalDistance)

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("The similarity score is: %d\n", similarityScore)
}
