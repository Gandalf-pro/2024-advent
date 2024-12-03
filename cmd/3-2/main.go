package main

import (
	"fmt"
	"regexp"

	"gandalf-pro.dev/advent-of-code-2024/pkg/api"
)

func main() {
	lines, _ := api.ReadLines("input.txt")

	total := 0

	include := true
	for _, line := range lines {
		re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
		muls := re.FindAllString(line, -1)
		for _, mul := range muls {
			if mul == "do()" {
				include = true
				continue
			}
			if mul == "don't()" {
				include = false
				continue
			}

			if !include {
				continue
			}

			var first int
			var last int

			fmt.Sscanf(mul, "mul(%d,%d)", &first, &last)
			total += first * last
		}
	}

	fmt.Println(total)
}
