package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if n == 0 {
		return false
	}
	binarySearch := func(m []int, t int, l int) bool {
		left, right := 0, l-1
		for left <= right {
			mid := m[(left+right)/2]
			if mid == t {
				return true
			}
			if mid < t {
				left = (left+right)/2 + 1
			} else {
				right = (left+right)/2 - 1
			}
		}
		return false
	}
	if m == 1 {
		return binarySearch(matrix[0], target, n)
	}
	local := -1
	tail := n - 1
	for left, right := 0, m-1; left <= right; {
		mid := (left + right) / 2
		if matrix[mid][tail] == target {
			return true
		}
		if matrix[mid][tail] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if left > right {
			local = left
		} else {
			local = mid
		}
	}
	if n == 1 || local > m-1 {
		return false
	}
	return binarySearch(matrix[local], target, n)

}
