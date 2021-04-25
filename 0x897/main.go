package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func increasingBST(root *TreeNode) *TreeNode {
//	stack := []*TreeNode{}
//	p := root
//	for p != nil {
//		if p.Left != nil {
//			stack = append(stack, p)
//			p = p.Left
//			continue
//		} else if p.Right != nil {
//			root = p
//			root = root.Right
//			stack = append(stack, p.Right)
//			p = p.Right
//			continue
//		} else {
//			p = stack[len(stack)-1]
//			stack = stack[:len(stack)-1]
//
//			fmt.Printf("p:%v\nstack:%v\n", p, stack)
//		}
//	}
//	return root
//}

func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode

	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		resNode.Right = node
		node.Left = nil
		resNode = node

		inorder(node.Right)
	}
	inorder(root)

	return dummyNode.Right
}
