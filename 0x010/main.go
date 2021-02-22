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
		temp := strings.Split(input.Text(), " ")
		fmt.Println(isMatch(temp[0], temp[1]))
	}
}

func isMatch(s string, p string) (aws bool) {
	slen, plen := len(s), len(p)
	if slen == 0 || plen == 0 {
		aws = false
	}
	point := 0
	for i := 0; i < plen; i++ {
		if p[i:i+1] == "." {

		}
	}
	return
}
