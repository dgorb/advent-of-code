package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dgorb/advent-of-code/2023/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	part1 := calculateValue(lines)
	part2 := calculateValue(replaceStringsWithDigits(lines))

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func calculateValue(lines []string) int {
	sum := 0
	for _, line := range lines {
		ints := regexp.MustCompile(`\d`).FindAllString(line, -1)
		if len(ints) > 0 {
			val, _ := strconv.Atoi(ints[0] + ints[len(ints)-1])
			sum += val
		}
	}
	return sum
}

func replaceStringsWithDigits(lines []string) []string {
	replacementTable := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4",
		"five":  "f5e",
		"six":   "s6",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	replacedLines := []string{}
	for _, line := range lines {
		replacedLine := line
		for k, v := range replacementTable {
			replacedLine = regexp.MustCompile(k).ReplaceAllString(replacedLine, v)
		}
		replacedLines = append(replacedLines, replacedLine)
	}
	return replacedLines
}
