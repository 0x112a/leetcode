package _x701

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode,val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	var insert func(father *TreeNode,node *TreeNode,int2 int)
	insert = func(father *TreeNode,node *TreeNode, int2 int) {
		if node == nil {
			if father.Val > val{
				father.Left = &TreeNode{Val: val}
				}else {
				father.Right = &TreeNode{Val: val}
			}
			return
		}
		if node.Val > val{
			insert(node,node.Left,val)
		}else {
			insert(node,node.Right,val)
		}
	}
	insert(root,root,val)

	return root
}