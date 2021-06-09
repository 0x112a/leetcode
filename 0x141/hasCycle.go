package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowP := head
	fastP := head.Next
	for fastP != nil {
		if fastP == slowP || fastP.Next == slowP {
			return true
		} else if fastP.Next != nil {
			fastP = fastP.Next.Next
		} else {
			fastP = fastP.Next
		}
		slowP = slowP.Next
	}
	return false
}
