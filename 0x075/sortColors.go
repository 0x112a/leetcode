package _x075

func sortColors(nums []int) {
	Left,Right := 0,len(nums)-1

	for i:=0; i<=Right; i++{
		for ; i<=Right && nums[i] == 2;Right--{
			nums[i],nums[Right] = nums[Right],nums[i]
		}

		if nums[i] == 0 {
			nums[i],nums[Left] = nums[Left],nums[i]
			Left++
		}
	}

}