package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	l := len(intervals)
	if l == 0 {
		return append(intervals, newInterval)
	}
	left, right := 0, l
	for i := 0; i < l; i++ {
		if newInterval[0] > intervals[i][1] {
			ans = append(ans, intervals[i])
			left++
		} else if newInterval[1] < intervals[i][0] {
			ans = append(ans, intervals[i])
			right--
		}
	}
	if left == l || right == 0 {
		ans = append(ans, newInterval)
		goto out
	}
	if newInterval[0] < intervals[left][0] {
		left = newInterval[0]
	} else {
		left = intervals[left][0]
	}
	if newInterval[1] > intervals[right-1][1] {
		right = newInterval[1]
	} else {
		right = intervals[right-1][1]
	}
	ans = append(ans, []int{left, right})
out:
	sort.Slice(ans, func(i, j int) bool { return ans[i][0] < ans[j][0] })
	return ans
}
