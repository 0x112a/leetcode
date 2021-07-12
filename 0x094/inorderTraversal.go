package _x094

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var	ans []int

	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil{
			return
		}
		inorder(node.Left)
		ans = append(ans,node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return ans
}