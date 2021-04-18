package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeDuplicates(nums []int) (answer int) {
	length := len(nums)
	if length == 0 || length == 1 {
		return length
	}
	answer = 1
	temp := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] == temp {
			continue
		} else {
			temp = nums[i]
			nums[answer] = temp
			answer++
		}
	}
	return answer

}
func rremoveDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	fast := 1
	temp := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] == temp {
			continue
		} else {
			temp = nums[i]
			nums[fast] = temp
			fast++
		}
	}
	return fast
}
