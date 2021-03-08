package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func findSubstring(s string, words []string) []int {
	l := len(s)
	t := len(words)
	var answer []int
	var stepSize int
	w := make(map[string]int, t)
	for _, t := range words {
		stepSize += len(t)
		w[t] = 0
	}

	return answer
}
