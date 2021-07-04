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
// 逆双指针
func Merge(nums1 []int,m int,nums2 []int, n int)  {
	p1,p2 := m-1,n-1
	cur := m+n-1
	for p1 >= 0 || p2 >=0 {
		if p1 < 0 {
			nums1[cur] = nums2[p2]
			p2--
		}else if p2 < 0{
			nums1[cur] = nums1[p1]
			p1--
		}else if nums1[p1] > nums2[p2] {
			nums1[cur] = nums1[p1]
			p1--
		}else if nums1[p1] <= nums2[p2]{
			nums1[cur] = nums2[p2]
			p2--
		}
		cur--
	}
}