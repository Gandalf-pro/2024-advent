package main

import (
	"fmt"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

func getValueAtCoordinates(lines []string, x int, y int) byte {
	if x < 0 || y < 0 || x >= len(lines[0]) || y >= len(lines) {
		return '.'
	}
	return lines[y][x]
}

func checkIfCoordinatedMatch(lines []string, coordinates [8]int) bool {
	str := "XMAS"
	for i := 0; i < 7; i += 2 {
		if getValueAtCoordinates(lines, coordinates[i], coordinates[i+1]) != str[i/2] {
			return false
		}
	}
	return true
}

func main() {
	lines, _ := api.ReadLines("input.txt")

	total := 0

	for y, line := range lines {
		for x, ch := range line {
			if ch != 'X' {
				continue
			}
			// Check top
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x, y - 1, x, y - 2, x, y - 3}) {
				total++
			}
			// Check diagonal top right
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x + 1, y - 1, x + 2, y - 2, x + 3, y - 3}) {
				total++
			}
			// Check right
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x + 1, y, x + 2, y, x + 3, y}) {
				total++
			}
			// Check diagonal bottom right
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x + 1, y + 1, x + 2, y + 2, x + 3, y + 3}) {
				total++
			}
			// Check bottom
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x, y + 1, x, y + 2, x, y + 3}) {
				total++
			}
			// Check diagonal bottom left
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x - 1, y + 1, x - 2, y + 2, x - 3, y + 3}) {
				total++
			}
			// Check left
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x - 1, y, x - 2, y, x - 3, y}) {
				total++
			}
			// Check diagonal top left
			if checkIfCoordinatedMatch(lines, [8]int{x, y, x - 1, y - 1, x - 2, y - 2, x - 3, y - 3}) {
				total++
			}
		}
	}

	fmt.Println(total)
}
