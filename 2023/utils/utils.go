package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func FileToStringSlice(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := []string{}
	lines = append(lines, strings.Split(string(content), "\n")...)

	return lines
}

func MaxInt(slice []int) int {
	slices.Sort(slice)
	return slice[len(slice)-1]
}
