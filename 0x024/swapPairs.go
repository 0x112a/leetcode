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
	//头结点总是指向第二的结点
	result = head.Next
	preP, P := head, head.Next
	// 记录前驱
	var temp *ListNode

	for P != nil {

		//交换指针
		preP.Next, P.Next = P.Next, preP

		//链接前驱
		if temp == nil {

		} else {
			temp.Next = P
		}

		//保存前驱节点指针
		temp = preP

		if preP.Next == nil {
			break
		}
		//移动双指针
		preP, P = preP.Next, preP.Next.Next

	}

	if result.Next == nil {
		head.Next = nil
		result.Next = head
	}

	return result
}
