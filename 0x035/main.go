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
		temp := input.Text()
		t := strings.Split(temp, ",")
		var a []int
		for i := 0; i < len(t); i++ {
			temp, _ := strconv.Atoi(t[i])
			a = append(a, temp)
		}
		fmt.Println(searchInsert(a, 7))
	}

}

func searchInsert(nums []int, target int) int {
	ans := len(nums)
	left, right := 0, ans-1
	for left <= right {
		mid := (right-left)>>1 + left
		fmt.Println(mid)
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
