package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")
	safeCount := 0
 	dampenedSafeCount := 0

	for _, l := range lines {
		report := strings.Split(l, " ")

		if isSafeReport(report) {
			safeCount++
			dampenedSafeCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			dampened := append([]string(nil), report[:i]...)
			dampened = append(dampened, report[i+1:]...)
			if isSafeReport(dampened) {
				dampenedSafeCount++
				break
			}
		}
	}

	fmt.Println("Part one:", safeCount)
	fmt.Println("Part two :", dampenedSafeCount)
}

func isSafeReport(report []string) bool {
    increaseCount, decreaseCount := 0, 0

    for i := 1; i < len(report); i++ {
        curr, _ := strconv.Atoi(report[i])
        prev, _ := strconv.Atoi(report[i-1])

        if utils.AbsInt(curr - prev) > 3 {
            return false
        }

        if curr > prev {
            increaseCount++
        } else if curr < prev {
            decreaseCount++
        }
    }

    return increaseCount == len(report) - 1 || decreaseCount == len(report) - 1
}
