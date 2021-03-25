package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		temp := strings.Split(input.Text(), " ")
		nums := []int{}
		for _, v := range strings.Split(temp[0], ",") {
			t, _ := strconv.Atoi(string(v))
			nums = append(nums, t)
			target, _ := strconv.Atoi(temp[1])
			fmt.Println(fourSum(nums, target))
		}
	}
}
func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	ans := [][]int{}
	for i := 0; i < l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, l-1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				if target == temp {
					fmt.Println("[", i, ",", j, ",", left, ",", right, "]", "--", temp)
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					var t int
					for t = left + 1; t < right && nums[left] == nums[t]; t++ {
					}
					left = t
					for t = right - 1; t > left && nums[right] == nums[t]; t-- {
					}
					right = t
				} else if temp < target {
					var t int
					for t = left + 1; t < right && nums[left] == nums[t]; t++ {
					}
					left = t

				} else {
					var t int
					for t = right - 1; t > left && nums[right] == nums[t]; t-- {
					}
					right = t
				}
			}
		}
	}
	return ans

}
