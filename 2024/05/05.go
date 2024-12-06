package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	input := utils.FileToStringSlice("input.txt")

	rules := make(map[string][]string)
	updates := [][]string{}

	// Read stuff into their slices
	for _, line := range input {
		if orderingRules := strings.Split(line, "|"); len(orderingRules) == 2 {
			ruleOne := strings.TrimSpace(orderingRules[0])
			ruleTwo := strings.TrimSpace(orderingRules[1])
			if _, exists := rules[ruleOne]; !exists {
				rules[ruleOne] = []string{}
			}
			rules[ruleOne] = append(rules[ruleOne], ruleTwo)
		}

		if update := strings.Split(line, ","); len(update) > 1 {
			for i, s := range update {
				update[i] = strings.TrimSpace(s)
			}
			updates = append(updates, update)
		}
	}

	findMid := func(s []string) string {
		return s[(len(s) - 1) / 2]
	}

	partOne := 0
	partTwo := 0

	for _, update := range updates {
		valid := true
		invalidUpdate := append([]string(nil), update...)

		for i, num := range update {
			if rule, exists := rules[num]; exists && utils.HasIntersection(rule, update[:i]) {
				valid = false

				for a, invNum := range invalidUpdate {
					if rulesForInvalidNum := rules[num]; a < i {
						for _, rule := range rulesForInvalidNum {
							if invNum == rule {
								invalidUpdate[a], invalidUpdate[i] = invalidUpdate[i], invalidUpdate[a]
								break
							}
						}
					}
				}
			}
		}

		processUpdate := func(upd []string) int {
			if mid, err := strconv.Atoi(findMid(upd)); err == nil {
				return mid
			}
			return 0
		}

		if valid {
			partOne += processUpdate(update)
		} else {
			partTwo += processUpdate(invalidUpdate)
		}
	}

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
