package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func removeElement(nums []int, val int) int {
	l := len(nums)
	p2 := l - 1
	for i := 0; i <= p2; i++ {
		if nums[i] == val {
			nums[p2], nums[i] = nums[i], nums[p2]
			i--
			p2--
		}
	}
	return p2 + 1
}
