package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func nthUglyNumber(n int) int {
	var a = []int{2, 3, 5}
	ugly := make([]int, n+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	p5, p2, p3 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := ugly[p2]*2, ugly[p3]*3, ugly[p5]*5
		ugly[i] = min(x2, min(x3, x5))
		if ugly[i] == x2 {
			p2++
		} else if ugly[i] == x3 {
			p3++
		} else {
			p5++
		}
	}
	return ugly[n]
}
