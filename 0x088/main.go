package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1)
	for i := 0; m < l; {
		nums1[m] = nums2[i]
		m++
		i++
	}
	sort.Ints(nums1)
}
