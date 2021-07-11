package _x083

type ListNode struct {
	Val int
	Next *ListNode
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
