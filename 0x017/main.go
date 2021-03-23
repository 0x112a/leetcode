package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(letterCombinations(input.Text()))
	}
}

var alpht map[string]([]string) = map[string]([]string){
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}
var ans []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans = []string{}
	backtrack(digits, 0, "")
	return ans
}
func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		ans = append(ans, combination)
	} else {
		digit := string(digits[index])
		letters := alpht[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+letters[i])
		}
	}
}
