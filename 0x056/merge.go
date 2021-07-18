package main

import (
	"sort"
)
//[[1,3],[2,6],[8,15],[10,18]]
//[[1,4],[4,5]]
//test sort
//[[1,4],[1,5]]
//[[1,4],[2,3]]
//[[2,3],[4,5],[6,7],[8,9],[1,10]]
func merge1(intervals [][]int)(ans [][]int) {
	if len(intervals) == 0|| len(intervals) ==1{
		return intervals
	}
	//sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]{
			return intervals[i][1] < intervals[j][1]
		}
		if intervals[i][0] > intervals[j][0]{
			return false
		}
		return true
	})
	slow,fast := 0,1

	for slow < len(intervals) {
		pre := intervals[slow][0]
		post := intervals[slow][1]
		for fast<len(intervals){
			if !(post < intervals[fast][0]) {
				slow++
				fast++
				if intervals[slow][1] > post{
					post = intervals[slow][1]
				}
			}else {
				break
			}
		}

		ans = append(ans,[]int{pre,post})
		slow++
		fast++
	}


	return
}
