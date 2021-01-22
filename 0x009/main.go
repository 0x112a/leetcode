package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		x, _ := strconv.Atoi(input.Text())
		fmt.Println(isPalindrome1(x))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	r := x
	sum := 0
	for r != 0 {
		sum = sum*10 + r%10
		r /= 10
	}
	if sum != x {
		return false
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	sum := 0
	for x > sum {
		sum = sum*10 + x%10
		x /= 10
	}
	return x == sum || x == sum/10
}
