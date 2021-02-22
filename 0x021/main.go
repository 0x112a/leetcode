package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("vim-go")

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		}
		return l2
	}
	if l2 == nil {
		return l1
	}
	var p1, p2, presult *ListNode
	p1, p2 = l1, l2
	if p1.Val > p2.Val {
		presult = p2
		p2 = p2.Next
	} else {
		presult = p1
		p1 = p1.Next
	}
	result := presult
	for p1 != nil {
		if p2 != nil && p1.Val > p2.Val {
			presult.Next = p2
			presult = presult.Next
			p2 = p2.Next
		} else {
			presult.Next = p1
			presult = presult.Next
			p1 = p1.Next
		}
	}
	if p2 != nil {
		presult.Next = p2
	}
	return result

}
