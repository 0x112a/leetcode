package main

func mypow(x float64, n int) float64 {

	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1.0 / quickMul(x, -n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	y := quickMul(x, n/2)

	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}
