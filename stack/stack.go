//定义链表
type ListNode struct {
    Val int
    Next *ListNode
}

//定义栈：栈中的元素使用链表实现
type Stack struct {
    Head, Bottom *ListNode
}

//新增元素
func (this *Stack) Push(x int)  {
    pnew:=new(ListNode)
    pnew.Val=x
    pnew.Next=this.Head
    this.Head=pnew
}

//弹出栈顶元素
func (this *Stack) Pop()  {
    this.Head=this.Head.Next    
}

//返回栈顶元素值
func (this *Stack) Top() int {
    return this.Head.Val
}

//取出栈顶中最小元素
func (this *Stack) GetMin() int {
    var min, tmp int
    pnew:=this.Head
    min=pnew.Val
    for pnew!=nil {
        tmp=pnew.Val
        if tmp<min {
            min=tmp
        }
        pnew=pnew.Next
    }
    return min
}

func Traverse(pstack *Stack) {
    if isempty(pstack) {
            return
    }
    p:=pstack.Head
    for ; p!=pstack.Bottom; p=p.Next {
        fmt.Println("%d", p.Val)
    }
    fmt.Println()
}

func isempty(pstack *Stack) bool {
    if pstack.Head==pstack.Bottom {
            return true
    }
    return false
}
