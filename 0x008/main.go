package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(myAtoi(input.Text()))
	}
}
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	l := len(s)
	symbol := 1
	i := 0
	if l == 0 {
		return 0
	}
	if s[0:1] == "-" {
		symbol = -1
		i = 1
	} else if s[0:1] == "+" {
		i = 1
	}
	result := 0
	if i < l {
		now, err := strconv.Atoi(s[i : i+1])
		i++
		if err != nil {
			return 0
		}
		result = now
	}
	for ; i < l; i++ {
		now, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			break
		}
		result = result*10 + now
		if i >= 8 {
			if result*symbol > math.MaxInt32 {
				return math.MaxInt32
			}
			if result*symbol < math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	return result * symbol
}
