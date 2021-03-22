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
func threeSumClosest(nums []int, target int) (ans int) {
	nums = bubsort(nums)
	l := len(nums)
	local := -1
	left, right := 0, l-1
	for i := l / 2; left != right-1; {
		fmt.Println("i:", i)
		if nums[i] == target {
			local = i
			left = i - 1
			right = i + 1
			break
		} else if nums[i] < target {
			left = i
			i = (right + left) / 2
		} else {
			right = i
			i = (right + left) / 2
		}
	}
	n := 0
	if local != -1 {
		ans += nums[local]
		n++
	}
	ads := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}
	for n < 3 {
		fmt.Println("left", left, "right", right)
		if left >= 0 && right <= l-1 {
			if ads(nums[left]-target) < ads(nums[right]-target) {
				ans += nums[left]
				n++
				left--
			} else {
				ans += nums[right]
				n++
				right++
			}
		} else if left < 0 {
			ans += nums[right]
			n++
			right++
		} else {
			ans += nums[left]
			n++
			left--
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
