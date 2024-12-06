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

	findMid := func(sloice []string) string {
		return sloice[(len(sloice) - 1) / 2]
	}

	partOne := 0
	for _, update := range updates {
		valid := true
		for i, num := range update {
			if rule, exists := rules[num]; exists && utils.HasIntersection(rule, update[:i]) {
				valid = false
				break
			}
		}

		if valid {
			if mid, err := strconv.Atoi(findMid(update)); err == nil {
				partOne += mid
			}
		}
	}

	fmt.Println("Part one:", partOne)
}
