package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func errortrap(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	var down, distance, sum int
	down = 0
	for i := 1; i < l || down != (l-1); i++ {
		fmt.Println("down:", down, "i:", i, "sum:", sum)
		if i == l {
			down++
			i = down
			distance = 0
			continue
		}
		if height[down] > height[i] {
			distance++
			continue
		} else {
			temp := distance * height[down]

			for j := down + 1; j < i; j++ {
				temp -= height[j]
				fmt.Println("temp:", temp, "distance:", distance)
			}
			distance = 0
			down = i
			sum += temp
		}
	}
	return sum

}
func trap(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	var down, sum, secondaryTop int
	down = 0

	subtrap := func(left, right int) int {
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		distance := right - left - 1
		temp := distance * h
		for j := left + 1; j < right; j++ {
			temp -= height[j]
		}
		fmt.Println("sub-------temp:", temp, "down:", left, "i:", right)
		return temp

	}
	for i := 1; i < l || down <= (l-2); {
		fmt.Println("down:", down, "i:", i, "sum:", sum)
		if i == l {
			sum += subtrap(down, secondaryTop)
			down = secondaryTop
			i = down + 1
			continue
		}
		for secondaryTop = i; i < l && height[down] > height[i]; i++ {

			if height[secondaryTop] <= height[i] {
				secondaryTop = i
			}
		}
		if i < l && height[down] <= height[i] {
			temp := subtrap(down, i)
			down = i
			sum += temp
			i++
		}

	}
	return sum

}
