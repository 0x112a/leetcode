package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func findMin(nums []int) int {
	l := len(nums)
	low, hight := 0, l-1
	for low < hight {
		mid := (hight + low) / 2
		if nums[mid] < nums[hight] {
			hight = mid - 1
		} else if nums[mid] > nums[hight] {
			low = mid + 1
		} else {
			hight--
		}
	}
	return nums[low]
}
