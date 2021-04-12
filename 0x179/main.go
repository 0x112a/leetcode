package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
}
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		x1, y1 := 10, 10
		for x1 <= x {
			x1 *= 10
		}
		for y1 <= y {
			y1 *= 10
		}
		return x1*y+x < y1*x+y
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, v := range nums {
		ans = append(ans, strconv.Itoa(v)...)
	}
	return string(ans)
}
