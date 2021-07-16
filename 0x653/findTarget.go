package _x653

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode,k int) bool {

	var find func(node *TreeNode,t int)bool

	var hashset = map[int]int{}

	find = func(node *TreeNode, t int) bool {
		if node == nil {
			return false
		}
		if _,ok := hashset[t-node.Val];ok{
			return true
		}
		hashset[node.Val]=node.Val
		return find(node.Left,t) || find(node.Right,t)
	}

	return find(root,k)
}