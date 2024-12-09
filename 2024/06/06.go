package main

import (
	"fmt"
	"strings"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	input := utils.FileToStringSlice("input.txt")

	directions := map[string][]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	dirOrder := map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}

	currDir := "up"
	visited := make(map[string]bool)

	var currentY, currentX int
	for y, line := range input {
		if x := strings.IndexRune(line, '^'); x != -1 {
			currentY, currentX = y, x
			break
		}
	}

	for {
		dir := directions[currDir]
		nextY, nextX := currentY+dir[0], currentX+dir[1]

		if nextY < 0 || nextY >= len(input) || nextX < 0 || nextX >= len(input[0]) {
			break
		}

		nextChar := input[nextY][nextX]
		if nextChar == '#' {
			currDir = dirOrder[currDir]
			continue
		}

		key := fmt.Sprintf("%d_%d", currentY, currentX)
		visited[key] = true

		currentY, currentX = nextY, nextX
	}

	fmt.Println("Part one:", len(visited))
}
