package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	slowa := headA
	fastb := headB

	for slowa != nil {
		for fastb != nil {
			if fastb == slowa {
				return fastb
			}
			fastb = fastb.Next
		}

		fastb = headB
		slowa = slowa.Next
	}

	return nil

}
func GreatgetIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	list1, list2 := headA, headB
	for list1 != list2 {
		if list1 != nil {
			list1 = list1.Next
		} else {
			list1 = headB
		}
		if list2 != nil {
			list2 = list2.Next
		} else {
			list2 = headA
		}
	}
	return list1

}
