package _x102

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode)(ans [][]int) {

	var level func([]*TreeNode)

	var queen []*TreeNode
	queen = append(queen,root)


	level = func(node []*TreeNode) {
		if len(node) == 0{
			return
		}
		var temp []int
		var tqueen []*TreeNode

		for _,v := range node{
			if v == nil {
				continue
			}
			temp = append(temp,v.Val)
			tqueen = append(tqueen,v.Left,v.Right)
		}
		if len(temp) != 0{
			ans = append(ans,temp)
		}
		level(tqueen)

	}
	level(queen)
	return
}