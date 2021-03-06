package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()
		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}

		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq[:] {
			if f != 0 {
				return 0
			}
		}

		for i := 1; i < length; i++ {
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}

			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, n) == 1
}
