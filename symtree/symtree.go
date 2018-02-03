type TreeNode *struct{
	  Val int
	  Left *TreeNode
	  Right *TreeNode
}

func IsSymtree(root *TreeNode) bool {
	  return Ismirror(root, root)
}

//use recursive
func Ismirror(t1 *TreeNode, t2 *TreeNode) bool {
	  if t1==nil && t2==nil {
		    return true
	  } else if t1==nil || t2==nil {
		    return false
	  } else if t1.Val==t2.Val {
		    return Ismirror(t1.Left, t2.Right) && Ismirror(t1.Right, t2.Left)
	  }
}
