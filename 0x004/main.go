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
		n1, n2 := strings.Split(temp[0], ","), strings.Split(temp[1], ",")
		fmt.Println(findMedianSortedArrays(convert(n1), convert(n2)))

	}
}
func convert(t []string) (n []int) {
	for i := 0; i < len(t); i++ {
		temp, _ := strconv.Atoi(t[i])
		n = append(n, temp)
	}
	return
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	var result []int
	l := len(nums1) + len(nums2)
	for i := 0; i < l; i++ {
		if len(nums1)-1 >= p1 && len(nums2)-1 >= p2 {
			if nums1[p1] > nums2[p2] {
				result = append(result, nums2[p2])
				p2++
			} else {
				result = append(result, nums1[p1])
				p1++
			}
		} else if p1 == len(nums1) {
			result = append(result, nums2[p2])
			p2++
		} else {
			result = append(result, nums1[p1])
			p1++
		}
	}
	var r float64
	if l%2 == 0 {
		r = float64(result[l/2]+result[l/2-1]) / 2.0
	} else {
		r = float64(result[l/2])
	}

	return r
}
