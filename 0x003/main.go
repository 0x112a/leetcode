package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(lengthofLongestSubstring(input.Text()))
	}
}

func lengthofLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	longest, temp := "", ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(temp, s[i:i+1]) {
			if len(longest) < len(temp) {
				longest = temp
			}
			n := strings.Index(temp, s[i:i+1])
			i = i - len(temp) + n + 1
			temp = s[i : i+1]
		} else {
			temp += s[i : i+1]
		}
	}
	if len(longest) < len(temp) {
		longest = temp
	}
	return len(longest)
}
