type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Maxdepth(root *TreeNode) int {
    if root=nil {
		    return 0
    }
    //use recursive
	  leftdepth:=Maxdepth(root.Left)+1
	  rightdepth:=Maxdepth(root.Right)+1
	  return Max(leftdepth, rightdepth)
}

func Max(a,b int) int {
	  if a>=b {
		    return a
	  } else {
		    return b
	  }
}
