package _x020

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	var s = "{[}]()"
	b := IsValid(s)
	fmt.Println(b)
}