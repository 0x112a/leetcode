package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// error

//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return nil
//	}
//	var ans *ListNode
//	var p *ListNode
//
//	cur := head
//	for cur != nil {
//		if cur.Val == val {
//			cur = cur.Next
//			continue
//		}
//		if ans == nil {
//			ans = cur
//			p = cur
//		}
//		p.Next = cur
//		cur = cur.Next
//
//	}
//
//	return ans
//}

func removeElements(head *ListNode, val int) *ListNode {
	v := &ListNode{Val: 0, Next: head}
	prev := v
	next := v.Next
	for next != nil {
		if next.Val == val {
			prev.Next = next.Next
			next = next.Next
		} else {
			prev = next
			next = next.Next
		}
	}
	return v.Next
}
