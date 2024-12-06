package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row, col int
}

type State struct {
	pos Position
	dir int
}

var directions = []Position{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func parseInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return grid, nil
}

func simulateGuard(grid []string, startPos Position, startDir int) (map[Position]bool, bool) {
	rows := len(grid)
	cols := len(grid[0])
	guardPos := startPos
	guardDir := startDir

	visited := make(map[Position]bool)
	path := make(map[State]int)
	loopThreshold := 5

	for {
		visited[guardPos] = true

		state := State{guardPos, guardDir}
		path[state]++
		if path[state] == loopThreshold {
			return visited, true
		}

		next := Position{
			row: guardPos.row + directions[guardDir].row,
			col: guardPos.col + directions[guardDir].col,
		}

		if next.row < 0 || next.row >= rows || next.col < 0 || next.col >= cols {
			return visited, false
		}

		if grid[next.row][next.col] == '#' {
			guardDir = (guardDir + 1) % 4
		} else {
			guardPos = next
		}
	}
}

func findLooperObstacles(grid []string, startPos Position, startDir int) int {
	rows := len(grid)
	cols := len(grid[0])
	looperObstacles := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '.' || (r == startPos.row && c == startPos.col) {
				continue
			}

			gridCopy := make([]string, rows)
			for i := 0; i < rows; i++ {
				gridCopy[i] = grid[i]
			}
			gridCopy[r] = gridCopy[r][:c] + "#" + gridCopy[r][c+1:]

			_, isLoop := simulateGuard(gridCopy, startPos, startDir)
			if isLoop {
				looperObstacles++
			}
		}
	}

	return looperObstacles
}

func main() {
	filename := ".\\input6.txt"

	grid, err := parseInput(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var startPos Position
	var startDir int
	dirMap := map[rune]int{'^': 0, '>': 1, 'v': 2, '<': 3}
	for r, row := range grid {
		for c, char := range row {
			if char == '^' || char == '>' || char == 'v' || char == '<' {
				startPos = Position{r, c}
				startDir = dirMap[char]
				break
			}
		}
	}

	visited, _ := simulateGuard(grid, startPos, startDir)
	fmt.Println("Distinct positions visited:", len(visited))

	looperObstacles := findLooperObstacles(grid, startPos, startDir)
	fmt.Println("Loop-causing positions:", looperObstacles)
}
