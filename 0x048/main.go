package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func rotate(matrix [][]int) {
	l := len(matrix)
	temp := make([][]int, l)
	for i := 0; i < l; i++ {
		temp[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := l - 1; j >= 0; j-- {
			fmt.Println(temp)
			temp[i][l-j-1] = matrix[j][i]
		}
	}
	copy(matrix, temp)
}
