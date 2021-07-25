package _x334

import "math"

func increasingTriple(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	Min,postMin := math.MaxInt32,math.MaxInt32

	for _,v := range nums{
		if Min >= v{
			Min=v
		}else if postMin >= v{
			postMin = v
		}else {
			return true
		}
	}

	return false
	
}
