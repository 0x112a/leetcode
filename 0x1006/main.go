package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func clumsy(N int) (ans int) {
	stack := []int{N}
	signal := 0
	N--
	for N > 0 {
		switch signal % 4 {
		case 0:
			stack[len(stack)-1] *= N
		case 1:
			stack[len(stack)-1] /= N
		case 2:
			stack = append(stack, N)
		default:
			stack = append(stack, -N)
		}
		N--
		signal++
	}
	for _, v := range stack {
		ans += v
	}

	return
}
