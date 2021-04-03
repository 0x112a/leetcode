package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	got := longestCommonSubsequence("abcde", "acde")
	want := 3
	if got != want {
		t.Errorf("this is a test ,don`t be warry-'%q','%q'", got, want)
	}
}
