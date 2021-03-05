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
		temp, _ := strconv.Atoi(input.Text())
		fmt.Println(intToRoman(temp))
	}
}
func intToRoman(num int) (answer string) {
	var values = [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var symbols = [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i := 12; i >= 0; i-- {
		if num/values[i] > 0 {
			temp := num / values[i]
			num %= values[i]
			for j := 0; j < temp; j++ {
				answer += symbols[i]
			}
		}
	}
	return answer
}
