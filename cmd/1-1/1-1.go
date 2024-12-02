package main

import (
	"fmt"
	"slices"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

func getLineValue(line string) (int, int) {
	var first int
	var last int

	fmt.Sscanf(line, "%d   %d", &first, &last)

	return first, last
}

func main() {
	lines, _ := api.ReadLines("input-1.txt")

	var leftList []int
	var rightList []int

	for _, line := range lines {
		left, right := getLineValue(line)
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	slices.SortFunc(leftList, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(rightList, func(a, b int) int {
		return a - b
	})
	var sum = 0
	for i := 0; i < len(leftList); i++ {
		var res = leftList[i] - rightList[i]
		if res < 0 {
			res *= -1
		}
		sum += res
	}
	fmt.Println(sum)
}
