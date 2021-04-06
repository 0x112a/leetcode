package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	slow, fast := 2, 2
	for fast < l {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			fast++
			slow++
		}
		fast++
	}
	return slow

}
