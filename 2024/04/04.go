package main

import (
	"fmt"
	"strings"

	"github.com/dgorb/advent-of-code/2024/utils"
)

func main() {
	input := utils.FileToStringSlice("input.txt")

	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}

func partOne(lines []string) int {
	height := len(lines)
	width := len(strings.TrimSpace(lines[0]))
	count := 0

	checkDirection := func(x, y, dx, dy int) bool {
		for i, c := range "MAS" {
			nx, ny := x+(i+1)*dx, y+(i+1)*dy
			if ny < 0 || ny >= height || nx < 0 || nx >= width {
				return false
			}
			if lines[ny][nx] != byte(c) {
				return false
			}
		}
		return true
	}

	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				directions := [][2]int{
					{1, 0},  // Right
					{-1, 0}, // Left
					{0, 1},  // Down
					{0, -1}, // Up
					{1, 1},  // Diagonal down-right
					{-1, -1}, // Diagonal up-left
					{-1, 1},  // Diagonal down-left
					{1, -1},  // Diagonal up-right
				}

				for _, dir := range directions {
					dx, dy := dir[0], dir[1]
					if checkDirection(x, y, dx, dy) {
						count++
					}
				}
			}
		}
	}

	return count
}

func partTwo(lines []string) int {
	height := len(lines)
	width := len(lines[0])
	count := 0

	checkPatterns := func(tl, tr, bl, br byte) bool {
		return (tl == 'M' && tr == 'S' && bl == 'M' && br == 'S') ||
			(tl == 'S' && tr == 'M' && bl == 'S' && br == 'M') ||
			(tl == 'S' && tr == 'S' && bl == 'M' && br == 'M') ||
			(tl == 'M' && tr == 'M' && bl == 'S' && br == 'S')
	}

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if lines[y][x] == 'A' {
				tl := lines[y-1][x-1]
				tr := lines[y-1][x+1]
				bl := lines[y+1][x-1]
				br := lines[y+1][x+1]

				if checkPatterns(tl, tr, bl, br) {
					count++
				}
			}
		}
	}

	return count
}
