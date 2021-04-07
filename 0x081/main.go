package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	if target == nums[left] || target == nums[right] {
		return true
	} else if right == 0 {
		return false
	} else if target < nums[left] {
		for nums[right-1] < nums[right] {
			if nums[right-1] == target {
				return true
			}
			right--
		}
	} else {
		for nums[left+1] > nums[left] {
			if nums[left+1] == target {
				return true
			}
			left++
		}
	}
	return false

}
