package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)

	if l == 0 || l == 1 {
		return false
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j <= i+k && j < l; j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= t {
				return true
			}
		}
	}
	return false
}
