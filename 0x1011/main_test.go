package main

import "testing"

func TestMain(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	got := shipWithinDays(weights, D)
	want := 15

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
