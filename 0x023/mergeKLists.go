package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(list []*ListNode) *ListNode {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 && list[0] == nil {
		return nil
	}

	var result = new(ListNode)
	p := result.Next

	var check func(l []*ListNode) bool

	check = func(l []*ListNode) bool {
		for _, v := range l {
			if v != nil {
				return true
			}
		}
		return false
	}

	for check(list) {
		var min *ListNode
		var index int
		for i, P := range list {

			if P != nil && min == nil {
				min = P
				index = i
			} else if P != nil && P.Val < min.Val {
				min = P
				index = i
			}
		}
		list[index] = min.Next

		p = min
		p = p.Next

	}

	return result.Next

}
