package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	input := utils.FileToString("input.txt")

	process := func(input string) int {
		result := 0
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		for _, m := range re.FindAllStringSubmatch(input, -1) {
			numOne, _ := strconv.Atoi(m[1])
			numTwo, _ := strconv.Atoi(m[2])
			result += numOne * numTwo
		}
		return result
	}

	reFilter := regexp.MustCompile(`don't\(\).*?(do\(\)|$)`)
	filteredInput := reFilter.ReplaceAllString(input, "")

	fmt.Println("Part one:", process(input))
	fmt.Println("Part two:", process(filteredInput))
}
