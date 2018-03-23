//找出树中国最长的一条 值相同 的路径
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Longestunivalued(root *TreeNode) int {
	var res int
  //递归：从下至上搜索
	if root!=nil {
		dfs(root, res)
	}
	return res
}

func dfs(node *TreeNode, ans int) int {
	var (
		l, r, resl, resr int
	)
	if node.Left!=nil {
		l=dfs(node.Left, ans)
	} else {
		l=0
	}
	if node.Right!=nil {
		r=dfs(node.Right, ans)
	} else {
		r=0
	}
	if node.Left!=nil && node.Val==node.Left.Val {
		resl=l+1
	} else {
		resl=0
	}
	if node.Right!=nil && node.Val==node.Right.Val {
		resr=r+1
	} else {
		resr=0
	}
  //出现的路径可能左右变换，需要保存此前路径相加的值
	if resl+resr>ans {
		ans=resl+resr
	}
	if resl>resr {
		return resl
	} else {
		return resr
	}
}
