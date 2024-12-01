package utils

import (
	"fmt"
	"os"
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

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
