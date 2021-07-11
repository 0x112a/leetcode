package _x206

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	ans := new(ListNode)

	p := head
	var q *ListNode
	for p != nil{
		q = p.Next
		p.Next = ans.Next
		ans.Next = p
		p = q
	}
	return ans.Next
}