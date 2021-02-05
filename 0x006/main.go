package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temp := strings.Split(input.Text(), ",")
		t, _ := strconv.Atoi(temp[1])
		fmt.Println(convert(temp[0], t))
	}
}

//func convert(s string, numRows int) string {
//	l := len(s)
//	direction := 0
//	dp := make([][]string, numRows)
//	for i := 0; i < numRows; i++ {
//		dp[i] = make([]string, l)
//	}
//	for j := 0; j < l; j++ {
//		if direction == 0 {
//			dp[j/numRows][j%numRows] = string(s[j])
//		} else {
//			dp[j/numRows]
//		}
//	}
//}

func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 || l <= numRows {
		return s
	}
	index := 2*numRows - 2
	result := make([]string, l)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < l; j += index {
			result = append(result, s[j+i:j+i+1])
			if i != 0 && i != numRows-1 && j+index-i < l {
				result = append(result, s[j+index-i:j+index-i+1])
			}
		}
	}
	return strings.Join(result, "")
}
