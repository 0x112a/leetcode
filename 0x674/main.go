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
		v := strings.Split(input.Text(), ",")
		var a []int
		for i := 0; i < len(v); i++ {
			temp, _ := strconv.Atoi(v[i])
			a = append(a, temp)
		}
		fmt.Println(findLengthOfLCIS(a))

	}
}
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	Max, temp := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			temp++
		} else {
			temp = 1
		}
		if temp > Max {
			Max = temp
		}
	}
	return Max
}
