package main

import "testing"

func TestMyPow(t *testing.T) {
	x := 2.0000
	n := 10
	got := mypow(x, n)
	want := 1024.00000

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
