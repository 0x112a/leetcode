package main

func numDecodings(s string) int {
	l := len(s)
	d := make([]int, l+1)
	d[0] = 1
	for i := 1; i <= l; i++ {
		if s[i-1] != '0' {
			d[i] += d[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			d[i] += d[i-2]
		}

	}
	return d[l]
}
