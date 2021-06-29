package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	// TODO 1.chek List is enough k
	check := func(h *ListNode)bool {
		t := h
		for i:=0; i<k; i++{
			if t == nil{
				return false
			}
			t=t.Next
		}
		return true
	}
	var result,h *ListNode
	p := head
	for check(p) {
		//头插法就地逆置单链表
		if h == nil{
			h = new(ListNode)
		}
		pre,n := p,p.Next
		for i:=0;i<k; i++{
			pre.Next = h.Next
			h.Next = pre
			pre = n
			if n.Next != nil{
				n = n.Next
			}else {

			}

		}
		p.Next = pre
		if result == nil {
			result = h
		}
		h = p
		p = p.Next

	}
	return result.Next

}