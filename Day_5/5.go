package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		rules[from] = append(rules[from], to)
	}

	for scanner.Scan() {
		line := scanner.Text()
		update := []int{}
		for _, val := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(val)
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}

func isValidUpdate(update []int, rules map[int][]int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for from, tos := range rules {
		for _, to := range tos {
			if posFrom, okFrom := position[from]; okFrom {
				if posTo, okTo := position[to]; okTo {
					if posFrom >= posTo {
						return false
					}
				}
			}
		}
	}

	return true
}

func correctOrder(update []int, rules map[int][]int) []int {
	updateRules := make(map[int][]int)
	inUpdate := make(map[int]bool)

	for _, page := range update {
		inUpdate[page] = true
	}

	for from, tos := range rules {
		if inUpdate[from] {
			for _, to := range tos {
				if inUpdate[to] {
					updateRules[from] = append(updateRules[from], to)
				}
			}
		}
	}

	inDegree := make(map[int]int)
	graph := make(map[int][]int)

	for from, tos := range updateRules {
		for _, to := range tos {
			graph[from] = append(graph[from], to)
			inDegree[to]++
		}
	}

	for _, page := range update {
		if _, exists := inDegree[page]; !exists {
			inDegree[page] = 0
		}
		if _, exists := graph[page]; !exists {
			graph[page] = []int{}
		}
	}

	queue := []int{}
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func findMiddle(update []int) int {
	mid := len(update) / 2
	return update[mid]
}

func main() {
	rules, updates, err := parseInput(".\\input5.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	validMiddleSum := 0
	var invalidUpdates [][]int

	for _, update := range updates {
		if isValidUpdate(update, rules) {
			validMiddleSum += findMiddle(update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fmt.Println("InvalidUpdate:", invalidUpdates)
	correctedMiddleSum := 0

	for _, update := range invalidUpdates {
		corrected := correctOrder(update, rules)
		correctedMiddleSum += findMiddle(corrected)
	}

	fmt.Println("Sum of middle page numbers (corrected invalid updates):", correctedMiddleSum)
}
