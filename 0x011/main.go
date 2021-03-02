package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func maxArea(height []int) (answer int) {
	l := len(height)
	for j := 0; j < l; j++ {
		for i := j + 1; i < l; i++ {
			if height[j] > height[i] {
				temp := height[i] * (i - j)
				if temp > answer {
					answer = temp
				}
			} else {
				temp := height[j] * (i - j)
				if temp > answer {
					answer = temp
				}
			}
		}
	}
	return

}
