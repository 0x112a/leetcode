package _x435

import "sort"

func eraseOverlapInervals(intervals [][]int) int {

	var	n = len(intervals)

	if intervals == nil{
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0]{
			return false
		}
		return true
	})

	dp := make([]int,n)
	for i:= range dp{
		dp[i] = 1
	}
	for i := 1; i< n;i++ {
		for j := 0;j<i;j++{
			if intervals[i][j] <= intervals[i][0]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
	}
	//可变参数
	return n-max(dp...)
}

func max(a ...int) int {
	res := a[0]

	for _,v := range a[1:]{
		if v > res{
			res = v
		}
	}
	return res
}
