package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func longestCommonSubsequence(text1 string, text2 string) (ans int) {
	l1, l2 := len(text1), len(text2)
	dynamic := make([][]int, l1+1)
	for i := range dynamic {
		dynamic[i] = make([]int, l2+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, char1 := range text1 {
		for j, char2 := range text2 {
			fmt.Println(i, j)
			if char1 == char2 {
				dynamic[i+1][j+1] = dynamic[i][j] + 1
			} else {
				dynamic[i+1][j+1] = max(dynamic[i][j+1], dynamic[i+1][j])
			}
		}
	}
	return dynamic[l1][l2]

}
