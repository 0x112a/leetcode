package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type NumArray struct {
	snums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) (asum int) {

	for ; i <= j; i++ {
		asum += this.snums[i]
	}
	return
}
