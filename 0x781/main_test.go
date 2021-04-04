package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	got := numRabbits([]int{1, 1, 2})
	want := 5
	if got != want {
		t.Errorf("some error push,got:'%q',wang:'%q'", got, want)
	}
}
