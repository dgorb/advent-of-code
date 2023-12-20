package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dgorb/advent-of-code/2023/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	partOne := 0
	partTwo := 0

	for _, l := range lines {
		score := 0

		numbers := strings.Split(strings.Split(l, ":")[1], "|")
		cardNumbers := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		for _, n := range myNumbers {
			if slices.Contains(cardNumbers, n) {
				if score == 0 {
					score = 1
					continue
				}
				score = score * 2
			}
		}
		partOne += score
	}

	partTwo += processSubset(lines)

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo+len(lines))
}

func processSubset(lines []string) int {
	score := 0

	for i, l := range lines {
		numMatches := 0

		numbers := strings.Split(strings.Split(l, ":")[1], "|")
		cardNumbers := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		for _, n := range myNumbers {
			if slices.Contains(cardNumbers, n) {
				numMatches += 1
			}
		}

		if numMatches > 0 {
			score += numMatches
			score += processSubset(lines[i+1 : i+1+numMatches])
		}
	}

	return score
}
