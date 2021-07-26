package _x001

import "sort"

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	n := len(nums)
	small, big := 0, n-1

	for nums[small]+nums[big] != target {
		//for small < big{
		if nums[small]+nums[big] < target {
			small++
		} else {
			big--
		}
	}
	return []int{small, big}
}

func TowSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{i, p}
		}
		hashTable[v] = i
	}
	return nil
}
