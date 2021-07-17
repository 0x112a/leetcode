package _x136

func singleNumber(nums []int) int {
	var ans = 0
	for _,v := range nums{
		ans = ans ^ v
	}

	return ans
}