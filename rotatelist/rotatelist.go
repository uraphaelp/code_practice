func Rotateright(head *ListNode, k int) *ListNode {
	if head==nil || head.Next==nil {
		return head
	}
	dmy:=new(ListNode)
	dmy.Val=0
	dmy.Next=head
	fast:=dmy
	slow:=dmy
	var (
		i, j int
	)
	//Get the length && move fast to the last element 取链表长度
	for i=0; fast.Next!=nil; i++ {
		fast=fast.Next
	}
	//move slow to the first 'need to move' element 取需要旋转的元素
	for j=i-k%i; j>0; j-- {
		slow=slow.Next
	}
	//pay attention to the rotation 旋转方法
	fast.Next=dmy.Next
	dmy.Next=slow.Next
	slow.Next=nil
	return dmy.Next
}
