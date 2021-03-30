package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("vim-go")
}
func rotateRight(head *ListNode, k int) (ans *ListNode) {
	l := 0
	var tail *ListNode
	for p := head; p != nil; p = p.Next {
		l++
		tail = p
	}
	tail.Next = head
	local := l - k%l
	for p := head; local > 0; local-- {
		if local == 1 {
			ans = p.Next
			p.Next = nil

		}
	}
	return ans
}
