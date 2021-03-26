package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	if p.Next == nil {
		return &ListNode{}
	}
	distance := 1
	var preWillDelete *ListNode
	for preWillDelete = p; p != nil; p = p.Next {
		if distance-1 == n && p.Next != nil {
			preWillDelete = preWillDelete.Next
			distance--
			continue
		}
		distance++
	}
	if preWillDelete == head {
		head = head.Next
	} else {
		preWillDelete.Next = preWillDelete.Next.Next
	}
	return head
}
