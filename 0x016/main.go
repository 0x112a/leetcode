package main

import (
	"bufio"
	"fmt"
	"math"
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
func threeSumClosest(nums []int, target int) (ans int) {
	nums = bubsort(nums)
	l := len(nums)
	ans = math.MaxInt32
	abs := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}
	direction := func(current int) bool {
		if abs(current-target) < abs(ans-target) {
			ans = current
		}
		if current > target {
			return true
		}
		return false
	}
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, l-1
		for left < right {
			current := nums[i] + nums[left] + nums[right]
			if current == target {
				return current
			}
			if direction(current) {
				temp := right - 1
				if left < temp && nums[temp] == nums[right] {
					temp--
				}
				right = temp
			} else {
				temp := left + 1
				if right > temp && nums[temp] == nums[left] {
					temp++
				}
				left = temp
			}
		}
	}
	return ans
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
