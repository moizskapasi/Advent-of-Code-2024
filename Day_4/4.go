package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countWord(grid [][]rune, word string) int {
	directions := [][2]int{
		{0, 1}, {0, -1},
		{1, 0}, {-1, 0},
		{1, 1}, {-1, -1},
		{1, -1}, {-1, 1},
	}

	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	count := 0

	isWordAt := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLength; i++ {
			newR := r + i*dr
			newC := c + i*dc
			if newR < 0 || newR >= rows || newC < 0 || newC >= cols || grid[newR][newC] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]
				if isWordAt(i, j, dr, dc) {
					count++
				}
			}
		}
	}

	return count
}

func countPattern(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' {
				if grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S' &&
					grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M' {
					count++
				}
				if grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M' &&
					grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S' {
					count++
				}
				if grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S' &&
					grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S' {
					count++
				}
				if grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M' &&
					grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M' {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	filename := "C:\\Users\\moizs\\Downloads\\AoC\\Day_4\\input4.txt"
	grid, err := parseInput(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	word := "XMAS"
	fmt.Println("Occurrences of XMAS:", countWord(grid, word))
	fmt.Println("Occurrences of X-MAS Pattern:", countPattern(grid))
}
