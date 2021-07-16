 package _x235
type TreeNode struct {
     Val   int
     Left  *TreeNode
     Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
     ancestor = root

     for{
          if p.Val < ancestor.Val && q.Val < ancestor.Val{
               ancestor = ancestor.Left
          }else if p.Val > ancestor.Val && q.Val > ancestor.Val {
               ancestor = ancestor.Right
          }else {
               return
          }
     }

}