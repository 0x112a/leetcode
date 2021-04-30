package main

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for num, t := range m {
		if t == 1 {
			return num
		}
	}
	return 0
}
