package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/dgorb/advent-of-code/2023/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	partOne := 0
	partTwo := len(lines)

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

	fmt.Println("Part one:", partOne)

	score := 0
	var wg sync.WaitGroup

	wg.Add(1)
	go processSubsets(lines, &score, &wg)

	wg.Wait()
	partTwo += score

	fmt.Println("Part two:", partTwo)

}

var mutex sync.Mutex

func processSubsets(lines []string, score *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, l := range lines {
		numMatches := 0

		numbers := strings.Split(strings.Split(l, ":")[1], "|")
		cardNumbers := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		for _, n := range myNumbers {
			if slices.Contains(cardNumbers, n) {
				numMatches++
			}
		}

		if numMatches > 0 {
			mutex.Lock()
			*score += numMatches
			mutex.Unlock()

			wg.Add(1)
			go processSubsets(lines[i+1:i+1+numMatches], score, wg)
		}
	}
}
