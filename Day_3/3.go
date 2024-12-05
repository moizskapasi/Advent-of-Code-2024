package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseInput(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return string(content)
}

func main() {
	filePath := ".\\input3.txt"
	input := parseInput(filePath)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	fmt.Printf("%v", matches)
	total := 0
	isEnabled := true

	for _, match := range matches {
		if len(match[0]) == 0 {
			continue
		}

		switch {
		case match[0] == "do()":
			isEnabled = true
		case match[0] == "don't()":
			isEnabled = false
		case match[1] != "" && match[2] != "":
			if isEnabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				total += num1 * num2
			}
		}
	}

	fmt.Printf("Total sum: %d\n", total)
}
