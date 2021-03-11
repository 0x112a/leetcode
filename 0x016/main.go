package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temp := strings.Split(input.Text(), " ")
		t, _ := strconv.Atoi(temp[1])
		ntemp := strings.Split(temp[0], ",")
		var n []int
		for _, i := range ntemp {
			t, _ := strconv.Atoi(i)
			n = append(n, t)
		}
		fmt.Println(threeSumClosest(n, t))
	}
}
func threeSumClosest(nums []int, target int) int {
	fmt.Println(bubsort(nums))
	return 0
}

func bubsort(nums []int) (ans []int) {
	n := len(nums)
	for i := n; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
