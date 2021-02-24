package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temp := strings.Split(input.Text(), " ")
		fmt.Println(isMatch(temp[0], temp[1]))
	}
}

func isMatch(s string, p string) (aws bool) {
	i, j := len(s), len(p)
	f := make([][]bool, i)
	for a := range f {
		f[a] = make([]bool, j)
	}

	return
}

func match(s string, p string, f [][]bool) (answer bool) {
	if p[len(p)-1:] == "*" {

	}
	return
}
