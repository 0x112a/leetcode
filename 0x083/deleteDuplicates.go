package _x083


func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	slowP,fastP := head,head.Next

	for fastP != nil {
		if fastP.Val == slowP.Val{
			fastP = fastP.Next
			continue
		}else {
			slowP.Next=fastP
			fastP = fastP.Next
			slowP = slowP.Next
		}

	}
	slowP.Next = fastP

	return head
}


