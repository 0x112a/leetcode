package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temp := input.Text()
		fmt.Println(longestPalindrome1(temp))
	}
}

//func longestPalindrome(s string) string {
//	if len(s) == 1 {
//		return s
//	}
//	if len(s) == 2 {
//		if s[0] == s[1] {
//			return s
//		} else {
//			return string(s[0])
//		}
//	}
//    
//	return ""
//}

func longestPalindrome(s string) string{
    l := len(s)
    ans := ""
    dp := make([][]int,l)
    for i:=0;i<l;i++{
        dp[i] = make([]int,l)
    }
    for n :=0;n<l;n++{
        for i :=0;i+n<l;i++{
            j:= i + n
            if n == 0{
                dp[i][j]=1
            }else if n == 1{
                if s[i]==s[j]{
                    dp[i][j] = 1
                }
            }else{
                if s[i]==s[j]{
                    dp[i][j]=dp[i+1][j-1]
                }
            }
            if dp[i][j] > 0 && n+1>len(ans){
                ans = s[i:i+n+1]
            }
        }
    }
    return ans

}
func longestPalindrome1(s string) string {
    n := len(s)
    ans := ""
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    for l := 0; l < n; l++ {
        for i := 0; i + l < n; i++ {
            j := i + l
            if l == 0 {
                dp[i][j] = 1
            } else if l == 1 {
                if s[i] == s[j] {
                    dp[i][j] = 1
                }
            } else {
                if s[i] == s[j] {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            if dp[i][j] > 0 && l + 1 > len(ans) {
                ans = s[i:i+l+1]
            }
        }
    }
    return ans
}

