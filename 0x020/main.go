package _x020

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(isValid(input.Text()))
	}
}
func bisValid(s string) bool {
	l := len(s)
	if l/2 != 0 {
		return false
	}
	pointLeft := l/2 - 1
	pointRight := l / 2
	res := map[string]string{"(": ")", "[": "]", "{": "}"}
	if pointLeft >= 0 {
		fmt.Println(res[s[pointLeft:pointLeft+1]] == s[pointRight:pointRight+1])
		if res[s[pointLeft:pointLeft+1]] == s[pointRight:pointRight+1] {
			pointLeft--
			pointRight--
		} else {
			return false
		}
	}
	return true
}
func isValid(s string) (answer bool) {
	l := len(s)
	if l%2 != 0 {
		return answer
	}
	res := map[string]string{"(": ")", "[": "]", "{": "}"}
	stack := []string{}
	top := -1
	for i := 0; i < l; i++ {
		if top == -1 {
			stack = append(stack, s[i:i+1])
			top++
		} else if res[stack[top]] == s[i:i+1] {
			top--
			if top == -1 {
				stack = []string{}
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i:i+1])
			top++
		}
	}
	if top == -1 {
		answer = true
	} else {
		answer = false
	}
	return answer
}
