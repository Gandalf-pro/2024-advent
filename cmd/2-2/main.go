package main

import (
	"fmt"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

func checkLine(numbers []int, skip int) bool {

	var lastValue int
	var isIncreasing bool
	i := 1

	if skip == 0 {
		lastValue = numbers[1]
		isIncreasing = numbers[1] < numbers[2]
		i = 2
	} else {
		lastValue = numbers[0]
		if skip == 1 {
			isIncreasing = numbers[0] < numbers[2]
		} else {
			isIncreasing = numbers[0] < numbers[1]
		}
	}

	for ; i < len(numbers); i++ {
		if skip == i {
			continue
		}
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

func getLineValue(line string) []int {
	numbers := make([]int, 0, 6)
	for i := 0; i < len(line); i++ {
		var tmp int
		fmt.Sscanf(line[i:], "%d", &tmp)
		numbers = append(numbers, tmp)
		i += api.GetDigitCount(tmp)
	}

	return numbers

}

func main() {
	lines, _ := api.ReadLines("input.txt")

	safeCount := 0

	var lineNumbers [][]int

	for _, line := range lines {
		lineNumbers = append(lineNumbers, getLineValue(line))
	}

	for _, numbers := range lineNumbers {
		foundTrue := checkLine(numbers, -1)
		for i := 0; i < len(numbers) && !foundTrue; i++ {
			foundTrue = checkLine(numbers, i)
		}
		if foundTrue {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
