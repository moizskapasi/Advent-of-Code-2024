package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func parseMap(filename string) (map[Point]rune, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	antennaMap := make(map[Point]rune)
	width, height := 0, 0

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, char := range line {
			if char != '.' {
				antennaMap[Point{x, y}] = char
			}
		}
		y++
	}
	height = y

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("error reading file: %v", err)
	}

	return antennaMap, width, height, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattanDistance(p Point) int {
	return abs(p.x) + abs(p.y)
}

func calculateAntinodesP1(antennaMap map[Point]rune, width, height int) []Point {
	groupedAntennas := make(map[rune][]Point)
	for pos, freq := range antennaMap {
		groupedAntennas[freq] = append(groupedAntennas[freq], pos)
	}

	antinodeSet := make(map[Point]bool)

	isInsideMap := func(p Point) bool {
		return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
	}

	for _, antennas := range groupedAntennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				p1 := antennas[i]
				p2 := antennas[j]

				dist1 := manhattanDistance(p1)
				dist2 := manhattanDistance(p2)

				var closer, farther Point
				if dist1 < dist2 {
					closer, farther = p1, p2
				} else {
					closer, farther = p2, p1
				}

				dx := farther.x - closer.x
				dy := farther.y - closer.y

				antinodeCloser := Point{closer.x + 2*dx, closer.y + 2*dy}
				antinodeFarther := Point{farther.x - 2*dx, farther.y - 2*dy}

				if isInsideMap(antinodeCloser) {
					antinodeSet[antinodeCloser] = true
				}
				if isInsideMap(antinodeFarther) {
					antinodeSet[antinodeFarther] = true
				}
			}
		}
	}

	var antinodePoints []Point
	for point := range antinodeSet {
		antinodePoints = append(antinodePoints, point)
	}

	return antinodePoints
}

func calculateAntinodesP2(antennaMap map[Point]rune, width, height int) []Point {
	groupedAntennas := make(map[rune][]Point)
	for pos, freq := range antennaMap {
		groupedAntennas[freq] = append(groupedAntennas[freq], pos)
	}

	antinodeSet := make(map[Point]bool)

	isInsideMap := func(p Point) bool {
		return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
	}

	for _, antennas := range groupedAntennas {
		if len(antennas) > 1 {
			for _, antenna := range antennas {
				antinodeSet[antenna] = true
			}

			for i := 0; i < len(antennas); i++ {
				for j := i + 1; j < len(antennas); j++ {
					p1 := antennas[i]
					p2 := antennas[j]

					dx := p2.x - p1.x
					dy := p2.y - p1.y

					gcd := abs(dx)
					if dy != 0 {
						for n := abs(dy); n != 0; {
							gcd, n = n, gcd%n
						}
					}
					dx /= gcd
					dy /= gcd

					for step := 1; ; step++ {
						newAntinode := Point{p1.x - step*dx, p1.y - step*dy}
						if isInsideMap(newAntinode) {
							antinodeSet[newAntinode] = true
						} else {
							break
						}
					}

					for step := 1; ; step++ {
						newAntinode := Point{p2.x + step*dx, p2.y + step*dy}
						if isInsideMap(newAntinode) {
							antinodeSet[newAntinode] = true
						} else {
							break
						}
					}
				}
			}
		}
	}

	var antinodePoints []Point
	for point := range antinodeSet {
		antinodePoints = append(antinodePoints, point)
	}

	return antinodePoints
}

func main() {
	filename := ".\\input8.txt"

	antennaMap, width, height, err := parseMap(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	antinodesPartOne := calculateAntinodesP1(antennaMap, width, height)
	fmt.Printf("Part One: Total unique antinodes: %d\n", len(antinodesPartOne))

	antinodesPartTwo := calculateAntinodesP2(antennaMap, width, height)
	fmt.Printf("Part Two: Total unique antinodes: %d\n", len(antinodesPartTwo))
}
