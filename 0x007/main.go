package main

import "fmt"
import "bufio"
import "os"
import "strconv"

//myself,High space complexity
func reverse(x int) int {
	var result int
	var stack []int
	for i := 0; x != 0; i++ {
		stack = append(stack, x%10)
		x /= 10
	}
	for i := 0; i < len(stack); i++ {
		result = result*10 + stack[i]

	}
	temp := int32(result)
	if int(temp) != result {
		return 0
	}
	return result
}

//other's optimal solution
func reverse1(x int) int {
	var res int
	if temp := int32(x); x != int(temp) {
		return 0
	}

	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if temp := int32(res); res != int(temp) {
		return 0
	}

	return res
}
func main() {
	//var s string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error:%v", err)
		}
		fmt.Println(reverse(n))
	}
}
