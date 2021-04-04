package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func numRabbits(answers []int) (sum int) {
	times := make(map[int]int)
	for _, v := range answers {
		times[v]++
	}
	fmt.Println(times)
	for i, v := range times {
		temp := v / (i + 1) * (i + 1)
		if v%(i+1) != 0 {
			temp += (i + 1)
		}
		sum += temp
	}
	return
}
