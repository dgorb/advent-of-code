package utils

import (
	"fmt"
	"os"
	"strings"
)

func FileToString(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.ReplaceAll(strings.ReplaceAll(string(content), "\r\n", ""), "\n", "")
}

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

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func HasIntersection(slice1, slice2 []string) bool {
	elements := make(map[string]struct{})

	for _, v := range slice1 {
		elements[v] = struct{}{}
	}

	for _, v := range slice2 {
		if _, exists := elements[v]; exists {
			return true
		}
	}

	return false
}
