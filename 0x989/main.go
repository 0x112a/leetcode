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
		var A []int
		s := strings.Split(input.Text(), ",")
		a := strings.Split(s[0], "")
		K, _ := strconv.Atoi(s[1])
		for i := 0; i < len(a); i++ {
			n, _ := strconv.Atoi(a[i])
			A = append(A, n)
		}
		fmt.Printf("%v\n", addToArrayForm(A, K))
	}
}
func addToArrayForm(A []int, K int) (ans []int) {
	for i := len(A) - 1; i >= 0 || K > 0; i-- {
		if i >= 0 {
			K += A[i]
		}
		ans = append(ans, K%10)
		K /= 10
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return

}
func MyaddToArrayForm(A []int, K int) []int {
	var result []int
	var sum int
	for i := 0; i < len(A); i++ {
		sum = sum*10 + A[i]

	}
	sum += K
	s := strconv.Itoa(sum)
	S := strings.Split(s, "")
	for i := 0; i < len(S); i++ {
		n, _ := strconv.Atoi(S[i])
		result = append(result, n)
	}
	return result
}
