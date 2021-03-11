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
		fmt.Println(searchInsert(n, t))
	}
}
func searchInsert(nums []int, target int) int {
	ans := len(nums)
	left, right := 0, ans-1
	for left <= right {
		mid := (right-left)>>1 + left //求得每次的中间位置坐标
		if target > nums[mid] {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}
	return ans

}
