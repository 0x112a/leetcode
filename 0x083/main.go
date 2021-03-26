package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val {
			p.Next = p.Next.Next

		} else {
			p = p.Next
		}
	}
	return head
}
