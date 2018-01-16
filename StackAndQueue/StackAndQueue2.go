//使用栈的结构实现队列
/* type ListNode struct {
    Val int
    Next *ListNode
}*/

type MyQueue struct {
    Top, Bottom *ListNode 
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    var queue MyQueue
    return queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    var tmp int
    var s2 MyQueue
    for !this.Empty() {
        pnew:=new(ListNode)
        tmp=this.Pop()
        pnew.Val=tmp
        pnew.Next=s2.Top
        s2.Top=pnew
    }
    node:=new(ListNode)
    node.Val=x
    node.Next=s2.Top
    s2.Top=node
    for !s2.Empty() {
        node_0:=new(ListNode)
        tmp=s2.Pop()
        node_0.Val=tmp
        node_0.Next=this.Top
        this.Top=node_0
    }
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    var val int
    if !this.Empty() {
        val=this.Top.Val
        this.Top=this.Top.Next
        return val
    }
    return 0
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if !this.Empty() {
        return this.Top.Val
    }
    return 0
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if this.Top==nil {
        return true
    }    
    return false
}
