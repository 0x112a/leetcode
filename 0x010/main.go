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

func isMatch(s string, p string) bool {
	i, j := len(s), len(p)
	matches := func(m, n int) bool {
		if m == 0 {
			return false
		}
		if p[n-1] == '.' {
			return true
		}
		return s[m-1] == p[n-1]
	}
	f := make([][]bool, i+1)
	for a := range f {
		f[a] = make([]bool, j+1)
	}
	f[0][0] = true
	for m := 0; m <= i; m++ {
		for n := 1; n <= j; n++ {
			if p[n-1] == '*' {
				f[m][n] = f[m][n] || f[m][n-2]
				if matches(m, n-1) {
					f[m][n] = f[m][n] || f[m-1][n]
				}
			} else if matches(m, n) {
				f[m][n] = f[m][n] || f[m-1][n-1]
			}
		}
	}
	return f[i][j]
}
