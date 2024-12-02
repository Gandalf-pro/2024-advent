package main

import (
	"fmt"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

func getLineValue(line string) bool {
	numbers := make([]int, 0, 6)
	for i := 0; i < len(line); i++ {
		var tmp int
		fmt.Sscanf(line[i:], "%d", &tmp)
		numbers = append(numbers, tmp)
		i += api.GetDigitCount(tmp)
	}

	lastValue := numbers[0]
	isIncreasing := lastValue < numbers[1]

	if lastValue == numbers[1] {
		return false
	}
	for i := 1; i < len(numbers); i++ {
		dif := numbers[i] - lastValue
		if dif < 0 {
			dif *= -1
		}
		if dif > 3 || dif < 1 {
			return false
		}
		if isIncreasing && numbers[i] <= lastValue {
			return false
		}
		if !isIncreasing && numbers[i] >= lastValue {
			return false
		}
		lastValue = numbers[i]
	}

	return true
}

func main() {
	lines, _ := api.ReadLines("input.txt")

	safeCount := 0

	for _, line := range lines {
		if getLineValue(line) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
