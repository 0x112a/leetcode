package _x053

func maxSubArray(nums []int) int {
	var i []int
	n := len(nums)
	i = make([]int,n+1)
	i[0] = nums[0]

	max := func(a,b int) int{
		if a > b{
			return a
		}
		return b
	}

	for index:=1; index < n; index++{
		i[index] = max(i[index-1]+nums[index],nums[index])
	}
// error Because i[n-1] is not max sub Array
	return i[n-1]
}

func MaxSubArray(nums []int) int {
	max := nums[0]


	for i,n :=1,len(nums); i< n ;i++{
		if nums[i-1]+nums[i]>nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max{
			max=nums[i]
		}
	}
	return max
}

