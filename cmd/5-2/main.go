package main

import (
	"fmt"
	"strconv"
	"strings"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

type OrderingRule struct {
	from int
	to   int
}

func checkIfLineIsOrderedCorrectly(values []int, orderingRules []OrderingRule) bool {
	for i, val := range values {
		for _, toVal := range values[i+1:] {
			for _, rule := range orderingRules {
				if rule.to == val && rule.from == toVal {
					return false
				}
			}

		}
	}
	return true
}

func orderOnce(values []int, orderingRules []OrderingRule) bool {
	for i, val := range values {
		for j, toVal := range values[i+1:] {
			for _, rule := range orderingRules {
				if rule.to == val && rule.from == toVal {
					// We swap and return false signifying not done
					values[i] = toVal
					values[j+i+1] = val
					return false
				}
			}

		}
	}
	return true
}

func orderLineAndReturnMiddle(values []int, orderingRules []OrderingRule) int {
	for !orderOnce(values, orderingRules) {
	}
	return values[len(values)/2]
}

func main() {
	lines, _ := api.ReadLines("input.txt")

	total := 0
	seenEmptyLine := false

	var orderingRules []OrderingRule

	for _, line := range lines {
		if len(line) == 0 {
			seenEmptyLine = true
			continue
		}
		if seenEmptyLine {
			values := make([]int, 0, 30)

			for _, valStr := range strings.Split(line, ",") {
				res, _ := strconv.Atoi(valStr)
				values = append(values, res)
			}
			if !checkIfLineIsOrderedCorrectly(values, orderingRules) {
				total += orderLineAndReturnMiddle(values, orderingRules)
			}
		} else {
			var rule OrderingRule
			fmt.Sscanf(line, "%d|%d", &rule.from, &rule.to)
			orderingRules = append(orderingRules, rule)
		}
	}

	fmt.Println(total)
}
