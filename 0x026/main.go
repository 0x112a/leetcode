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
