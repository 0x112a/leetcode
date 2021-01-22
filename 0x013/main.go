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
		s := input.Text()
		fmt.Println(romanToInt(s))
	}
}

func romanToInt(s string) int {
	S := strings.Split(s, "")
	var roman = make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000
	var result int
	for i := 0; i < len(S); i++ {
		if i != len(S)-1 {

			fmt.Println(i)
			if S[i] == "I" && (S[i+1] == "V" || S[i+1] == "X") {
				result += (roman[S[i+1]] - roman[S[i]])
				i++
				continue
			}
			if S[i] == "X" && (S[i+1] == "L" || S[i+1] == "C") {
				result += (roman[S[i+1]] - roman[S[i]])
				i++
				continue
			}
			if S[i] == "C" && (S[i+1] == "D" || S[i+1] == "M") {
				result += (roman[S[i+1]] - roman[S[i]])
				i++
				continue
			}
		}
		result += roman[S[i]]

	}
	return result
}
