package _x144

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// error
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	preorder(root,ans)
	return ans
}

func preorder(p *TreeNode, ans []int)  {
	if p==nil{
		return
	}
	ans = append(ans,p.Val)
	preorder(p.Left,ans)
	preorder(p.Right,ans)
}
//func preorderTraversal(root *TreeNode) []int {
//	var ans []int
//	var preorder func(*TreeNode)
//
//	preorder=func(p *TreeNode)  {
//		if p == nil {
//			return
//		}
//		ans = append(ans,p.Val)
//		preorder(p.Left)
//		preorder(p.Right)
//	}
//	preorder(root)
//	return ans
//}