package _x119

func getRow(rowIndex int)[]int {
	if rowIndex == 0{
		return nil
	}
	var get func([]int)[]int

	get = func(ints []int) (ans []int) {
		n := len(ints)
		for i:=0;i<n+1;i++{

			if i-1<0 {
				ans = append(ans,ints[i])
				continue
			}
			if i>=n {
				ans = append(ans,ints[i-1])
				continue
			}

			ans = append(ans,ints[i-1]+ints[i])
		}
		return
	}
	a := []int{1}
	for i:=0;i<rowIndex;i++{
		a=get(a)
	}
	return a
}