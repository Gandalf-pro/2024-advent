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
			var values []int
			for _, valStr := range strings.Split(line, ",") {
				res, _ := strconv.Atoi(valStr)
				values = append(values, res)
			}
			if checkIfLineIsOrderedCorrectly(values, orderingRules) {
				total += values[len(values)/2]
			}
		} else {
			var rule OrderingRule
			fmt.Sscanf(line, "%d|%d", &rule.from, &rule.to)
			orderingRules = append(orderingRules, rule)
		}
	}

	fmt.Println(total)
}
