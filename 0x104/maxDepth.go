package _x104


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (ans int) {
	var max func([]*TreeNode)

	max = func(node []*TreeNode) {
		if len(node) == 0 {
			return
		}
		var temp []*TreeNode
		for _,v := range node{
			if v != nil {
				temp = append(temp,v.Left,v.Right)
			}
		}
		ans++
		max(temp)
	}

	max([]*TreeNode{root})
	ans--
	return
}