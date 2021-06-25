package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d", &n)
	r := generateParenthesis(n)
	fmt.Println(r)
}

var result []string

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}

	getParenthesis("", n, n)

	return result
}

func getParenthesis(str string, left int, right int) {
	if left == 0 && right == 0 {
		result = append(result, str)
		return
	}

	if left == right {
		getParenthesis(str+"(", left-1, right)
	} else if left < right {
		if left > 0 {
			getParenthesis(str+"(", left-1, right)
		}
		getParenthesis(str+")", left, right-1)
	}
}
