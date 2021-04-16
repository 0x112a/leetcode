package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	dp := make([]int, l)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}
