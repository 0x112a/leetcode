package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}

}

func (this *BSTIterator) Next() int {
	for p := this.cur; p != nil; p = p.Left {
		this.stack = append(this.stack, p)
	}
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	ans := this.cur.Val
	this.cur = this.cur.Right
	return ans

}
func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}
