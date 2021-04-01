package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}
func divide(dividend int, divisor int) (ans int) {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return dividend * -1
		}
		return math.MaxInt32
	}
	if dividend < 0 {
		dividend *= -1
		if divisor < 0 {
			divisor *= -1
			ans = 1
		} else {
			ans = -1
		}
	} else {
		if divisor < 0 {
			divisor *= -1
			ans = -1
		} else {
			ans = 1
		}
	}
	div := func(a int, b int) int {
		count := 0
		temp := b
		targer := 0
		for a >= b {
			targer = 1
			if a < (b << 1) {
				a -= b
				b = temp
				if a < b {
					targer = 0
					break
				}
				continue
			}
			count++
			b <<= 1
		}
		if count == 0 {
			return targer
		}
		return (1 << count) + targer
	}

	//	count := 0
	//	for dividend >= divisor {
	//		count++
	//		dividend -= divisor
	//	}
	//	ans *= count
	return div(dividend, divisor) * ans

}
