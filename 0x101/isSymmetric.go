package _x101

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(*TreeNode,*TreeNode)bool

	check = func(node *TreeNode, node2 *TreeNode)bool {
		if node == nil && node2 == nil{
			return true
		}else if node == nil || node2 ==nil {
			return false
		}
		return node.Val == node2.Val && check(node.Left,node2.Right) && check(node.Right,node2.Left)
	}

	return check(root.Left,root.Right)
}