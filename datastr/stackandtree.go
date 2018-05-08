package tree_build

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//----------------------------以下是stack的定义和操作-----------------------------------
type Stack struct {
	Element []interface{}
}

func Newstack() *Stack {
	return &Stack{}
}

func (stack *Stack) Size() int {
	return len(stack.Element)
}

func (stack *Stack) Empty() bool {
	if stack.Element == nil || stack.Size() == 0 {
		return true
	}
	return false
}

func (stack *Stack) Print() {
	for i := len(stack.Element) - 1; i >= 0; i--{
		fmt.Println(i,"=>",stack.Element[i])
	}
}

func (stack *Stack) Top() interface{} {
	if stack.Size() > 0 {
		return stack.Element[stack.Size() - 1]
	}
	return nil //read empty stack
}

func (stack *Stack) Push(value ...interface{}) {
	stack.Element = append(stack.Element,value...)
}

func (stack *Stack) Pop() interface{} {
	if stack.Size() > 0 {
		res:=stack.Element[stack.Size()-1]
		stack.Element = stack.Element[:stack.Size() - 1]
		return res
	}
	return nil
}

//随机数构建随机层的二叉搜索树
func New(n int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(n) {
		//perm 产生 0---(n-1) 的一个随机排序切片: 所以构造节点值为任意值的搜索树时，可以将 Perm(n) 换成一个预先定义好的数组
		t = Insert(t, 1+v)
	}
	defer func() {
		Preorder(t)
		fmt.Printf("------\n")
		Inorder(t)
		fmt.Print("\n")
	}()
	return t
}

//递归才能保证最后返回的是根节点
func Insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if v < t.Val {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

//先序遍历
func Preorder(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	Preorder(t.Left)
	Preorder(t.Right)
}

//中序遍历
func Inorder(t *TreeNode) {
	if t == nil {
		return
	}
	Inorder(t.Left)
	fmt.Println(t.Val)
	Inorder(t.Right)
}

//镜像翻转
func Rotate(t *TreeNode) *TreeNode {
	if t==nil {
		return t
	}
	tmp:=t.Left
	t.Left=Rotate(t.Right)
	t.Right=Rotate(tmp)
	return t
}
//镜像翻转非递归：借助栈
func Mirror(t *TreeNode) {
	if t==nil {
		return
	}
	s:=Newstack()
	//s.Element=[]TreeNode{}
	s.Push(t)
	for !s.Empty() {
		tree, _:=s.Pop().(TreeNode)
		//类型断言
		if tree.Left!=nil || tree.Right!=nil {
			tmp:=tree.Left
			tree.Left=tree.Right
			tree.Right=tmp
		}
		if tree.Left!=nil {
			s.Push(tree.Left)
		}
		if tree.Right!=nil {
			s.Push(tree.Right)
		}
	}
}
