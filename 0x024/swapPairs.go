package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	} else if head.Next == nil {
		return head
	}

	var result *ListNode
	result = head.Next
	preP, P := head, head.Next

	for preP != nil && P.Next != nil {

		preP.Next = P.Next

		P.Next = preP

		preP = preP.Next

		P = preP.Next.Next

	}

	return result
}
