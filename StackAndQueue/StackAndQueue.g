type ListNode struct {
    Val int
    Next *ListNode
}

type MyStack struct {
    Front, Back *ListNode //使用队列的定义形式：有队头，有队尾    
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    var stack MyStack
    return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    var sz int
    pnew:=new(ListNode)
    pnew.Val=x
    pnew.Next=nil
    //注意这个地方，怎么连接一个新的节点进去
    if this.Empty() {
        this.Front=pnew
    } else {
        this.Back.Next=pnew
    }
    this.Back=pnew
    sz=this.Size()
    //把新插入的节点放在front
    for i:=0; i<sz-1; i++  {
        node:=new(ListNode)
        tmp:=this.Pop()
        node.Val=tmp
        node.Next=nil
        if this.Empty() {
            this.Front=node
        } else {
            this.Back.Next=node
        }
        this.Back=node
    }
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    var val int
    if !(this.Empty()) {
        val=this.Front.Val
        this.Front=this.Front.Next
    }
    return val
}


/** Get the top element. */
func (this *MyStack) Top() int {
    if !(this.Empty()) {
        return this.Front.Val
    }
    return 0
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if this.Front==nil {
        return true
    }
    return false
}

func (this *MyStack) Size() int {
    var i int
    var pnew *ListNode
    pnew=this.Front
    for pnew!=nil {
        i++
        pnew=pnew.Next
    }
    return i
}
