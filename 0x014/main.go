package main

import (
	"fmt"
	"go/src/bufio"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strs := strings.Split(input.Text(), " ")
		fmt.Println(llongestCommonPrefix(strs))

	}
}

func llongestCommonPrefix(strs []string) (result string) {
	if len(strs) == 0 {
		result = ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				//result = strs[0][:i]
				return strs[0][:i]
			}
		}
	}
	return
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
