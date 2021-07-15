package _x226

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		 return nil
	}

	var invert func(node *TreeNode)

	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left,node.Right = node.Right,node.Left
		invert(node.Left)
		invert(node.Right)

	}
	invert(root)

	return root
}