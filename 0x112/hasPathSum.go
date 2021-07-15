package _x112

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode,targetSum int) bool {
	if root == nil {
		return false
	}
	//var hasPath func(node *TreeNode,target int)bool
	//
	//hasPath = func(node *TreeNode, target int) bool {
	//	if node.Left == nil && node.Right == nil{
	//		return targetSum == root.Val
	//	}
	//}

	if root.Left == nil && root.Right==nil{
		return root.Val == targetSum
	}

	return hasPathSum(root.Left,targetSum-root.Val) || hasPathSum(root.Right,targetSum-root.Val)
}