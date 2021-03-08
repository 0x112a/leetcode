package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	if nlen == 0 {
		return 0
	}
	if hlen == 0 {
		return -1
	}
	ans := strings.Index(haystack, needle)
	return ans

}
