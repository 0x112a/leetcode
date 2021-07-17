package main

import "sort"

func threeSum(num []int) [][]int {
	sort.Ints(num)
	n := len(num)
	ans := make([][]int,0)

	for first := 0; first < n ; first++{
		if first > 0 && num[first-1] == num[first]{
			continue
		}
		target := -1 * num[first]
		Right := n-1

		for second := first+1; second < n; second++{
			if second > first+1 && num[second] == num[second-1]{
				continue
			}

			for second < Right && num[Right]+num[second] > target{
				Right--
			}

			if second == Right{
				break
			}
			if num[second]+num[Right]==target {
				ans = append(ans,[]int{num[first],num[second],num[Right]})
			}
		}

	}

	return ans
}
