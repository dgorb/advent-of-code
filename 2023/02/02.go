package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dgorb/advent-of-code/2023/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	partOne := 0
	partTwo := 0

	for _, l := range lines {
		gameId, err := strconv.Atoi(regexp.MustCompile(`Game (\d+)`).FindStringSubmatch(l)[1])
		if err != nil {
			os.Exit(1)
		}

		colorCounts := map[string][]int{}

		counts := regexp.MustCompile(`\d+\s(red|green|blue)`).FindAllString(l, -1)
		for _, c := range counts {
			parts := strings.Split(c, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				os.Exit(1)
			}
			color := parts[1]

			if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
				gameId = 0
			}
			colorCounts[color] = append(colorCounts[color], count)
		}

		partOne += gameId
		partTwo += (utils.MaxInt(colorCounts["red"]) * utils.MaxInt(colorCounts["green"]) * utils.MaxInt(colorCounts["blue"]))
	}

	fmt.Println("Part 1:", partOne)
	fmt.Println("Part 2:", partTwo)
}
