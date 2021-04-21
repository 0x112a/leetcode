package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	got := numDecodings("12")
	excepted := 2
	if got != excepted {
		t.Errorf("excepted '%d',but got '%d'", excepted, got)
	}
}
