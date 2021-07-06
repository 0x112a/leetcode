package _x350

import "sort"

//hash error
//func intersect(nums1 []int, nums2 []int) []int {
//	var setHash = map[int]int{}
//	for _,v :=  range nums1{
//		setHash[v]++
//	}
//	for _,v :=  range nums2{
//		setHash[v]++
//	}
//	var result []int
//	for k,v := range setHash{
//		t := v/2
//		for i:=0;i<t; i++{
//			result = append(result,k)
//		}
//	}
//	return result
//}
func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1,p2 := len(nums1)-1,len(nums2)-1
	for p1>=0 && p2 >=0{
		if nums1[p1] == nums2[p2]{
			result = append(result,nums1[p1])
			p1--
			p2--
		}else if nums1[p1] > nums2[p2] {
			p1--
		}else {
			p2--
		}
	}
	return result
}