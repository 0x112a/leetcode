package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func shipWithinDays(weights []int, D int) int {
	var leftMax, rightSum int
	for _, v := range weights {
		if v > leftMax {
			leftMax = v
		}
		rightSum += v
	}
	for leftMax < rightSum {
		mid := (leftMax + rightSum) / 2
		s := 0
		day := 1
		for _, v := range weights {
			if s+v > mid {
				day++
				s = 0
			}
			s += v
		}
		if day <= D {
			rightSum = mid
		} else {
			leftMax = mid + 1
		}
	}
	return leftMax
}
