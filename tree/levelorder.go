package tree

//树节点定义
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//队列节点定义
type QueueNode struct {
	Node *TreeNode
	Next *QueueNode
}

//队列定义
type Queue struct {
	Front *QueueNode
	Back *QueueNode
}

func Levelorder(root *TreeNode) [][]int {
	var result [][]int
	if root==nil {
		return result
	}
	queue:=new(Queue)
	Push(queue, root)
	for !Isempty(queue) {
		var (
			tmp []int
			sz = Size(queue)
		)
		for i:=0; i<sz; i++ {
			tmptree:=Peek(queue)
			tmp=append(tmp, tmptree.Val)
			if tmptree.Left!=nil {
				Push(queue, tmptree.Left)
			}
			if tmptree.Right!=nil {
				Push(queue, tmptree.Right)
			}
		}
		result=append(result, tmp)
	}
	return result
}

func Push(q *Queue, t *TreeNode)  {
	qnew:=new(QueueNode)
	qnew.Node=t
	qnew.Next=nil
	if Isempty(q) {
		q.Front=qnew
		q.Back=qnew
	} else {
		//注意这个地方的指针转换
		q.Back.Next=qnew
		q.Back=qnew
	}
}

func Peek(q *Queue) *TreeNode {
	node:=new(TreeNode)
	if Isempty(q){
		return node
	} else {
		node=q.Front.Node
		q.Front=q.Front.Next
	}
	return node
}

func Isempty(q *Queue) bool {
	if q.Front==nil {
		return true
	} else {
		return false
	}
}

func Size(q *Queue) int {
	var count int
	qnode:=new(QueueNode)
	qnode=q.Front
	for qnode!=nil {
		qnode=qnode.Next
		count++
	}
	return count
}
