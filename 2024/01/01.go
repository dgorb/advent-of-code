package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	partOne := 0
	partTwo := 0

	left := []int{}
	right := []int{}

	for _, l := range lines {
		split := strings.Split(l, "   ")

		l, err := strconv.Atoi(split[0])
		if err != nil {
			break
		}

		r, err := strconv.Atoi(split[1])
		if err != nil {
			break
		}

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		partOne += utils.AbsInt(left[i] - right[i])

		countInRight := 0
		for j := range right {
			if left[i] == right[j] {
				countInRight += 1
			}
		}
		partTwo += left[i] * countInRight
	}

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
