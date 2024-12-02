package main

import (
	"fmt"

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

	rightSideOccurrenceMap := make(map[int]int)
	for i := 0; i < len(rightList); i++ {
		rightSideOccurrenceMap[rightList[i]] = rightSideOccurrenceMap[rightList[i]] + 1
	}

	var sum = 0
	for i := 0; i < len(leftList); i++ {
		var occurred = rightSideOccurrenceMap[leftList[i]]
		sum += leftList[i] * occurred
	}

	fmt.Println(sum)
}
