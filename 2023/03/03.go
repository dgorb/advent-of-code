package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/dgorb/advent-of-code/2023/utils"
)

func main() {
	lines := utils.FileToStringSlice("input.txt")

	part1 := 0
	twoDSlice := to2DSlice(lines)

	for y := range twoDSlice {
		digit := []rune{}
		for x := range twoDSlice[y] {
			r := twoDSlice[y][x]
			if unicode.IsDigit(r) {
				digit = append(digit, r)
			} else if len(digit) > 0 {
				for i := range digit {
					if isAdjacent(twoDSlice, y, x-i-1) {
						val, _ := strconv.Atoi(string(digit))
						part1 += val
						break
					}
				}
				digit = []rune{}
			}
		}
	}

	fmt.Println("Part 1:", part1)
}

func to2DSlice(lines []string) [][]rune {
	var charSlice [][]rune
	for _, line := range lines {
		chars := []rune(line)
		charSlice = append(charSlice, chars)
	}
	return charSlice
}

func isAdjacent(slice [][]rune, y, x int) bool {
	isSymbol := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) && r != '.'
	}

	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			if dy == 0 && dx == 0 {
				continue
			}
			newY, newX := y+dy, x+dx
			if newY >= 0 && newY < len(slice) && newX >= 0 && newX < len(slice[0]) {
				if isSymbol(slice[newY][newX]) {
					return true
				}
			}
		}
	}
	return false
}
