package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
	merge()
}
func merge(intervals [][]int) (ans [][]int) {
	l := len(intervals)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	intervals = sort(intervals, l)
	pre := intervals[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < l; i++ {
		cur := intervals[i]
		if pre[1] < cur[0] {
			ans = append(ans, pre)
			pre = intervals[i]
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	ans = append(ans, pre)
	return
}

func sort1(intervals [][]int, l int) [][]int {
	for i := 0; i < l-1; i++ {
		if intervals[i][0] > intervals[i+1][0] {
			intervals[i][0], intervals[i+1][0] = intervals[i+1][0], intervals[i][0]
			intervals[i][1], intervals[i+1][1] = intervals[i+1][1], intervals[i][1]
		}
	}

	return intervals
}
