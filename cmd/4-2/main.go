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

func main() {
	lines, _ := api.ReadLines("input.txt")

	total := 0

	for y, line := range lines {
		for x, ch := range line {
			if ch != 'A' {
				continue
			}
			topLeft := getValueAtCoordinates(lines, x-1, y-1)
			topRight := getValueAtCoordinates(lines, x+1, y-1)
			bottomLeft := getValueAtCoordinates(lines, x-1, y+1)
			bottomRight := getValueAtCoordinates(lines, x+1, y+1)
			tmp := 0
			if topLeft == 'M' && bottomRight == 'S' {
				tmp++
			}
			if topLeft == 'S' && bottomRight == 'M' {
				tmp++
			}
			if topRight == 'S' && bottomLeft == 'M' {
				tmp++
			}
			if topRight == 'M' && bottomLeft == 'S' {
				tmp++
			}
			if tmp >= 2 {
				total++
			}
		}
	}

	fmt.Println(total)
}
