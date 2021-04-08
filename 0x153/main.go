package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func findMin(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	left, right := 0, l-1
	if nums[left] > nums[right] {
		for right > left && nums[right-1] <= nums[right] {
			right--
		}
		return nums[right]
	}
	return nums[left]
}
