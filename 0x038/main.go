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
		fmt.Println(countAndSay(strconv.Atoi(input.Text())))
	}
}
func countAndSay(n int, e error) string {
	ans := "1"
	for i := 1; i < n; i++ {
		now := ans[0:1]
		times := 0
		temp := ""
		for _, v := range ans {
			fmt.Println(i, "ans:", string(v), "times:", times)
			if now == string(v) {
				times++
			} else {
				temp += strconv.Itoa(times) + now
				now = string(v)
				times = 1
			}
		}
		temp += strconv.Itoa(times) + now
		ans = temp
	}
	return ans
}
