//给定链表：移除从尾部开始的第n个元素

func Removenth(head *List.ListNode, n int) *List.ListNode {
	this:=Findlastnth(head, n)
	if this==head {
		head=head.Next
	} else {
		last:=Findlastnth(head, n-1)
		last.Next=this.Next
	}
}

//找到第n个
func Findlastnth(head *List.ListNode, n int) *List.ListNode {
	var count int
	p:=head
  //注意循环范围
	for count<GetLength(head)-n {
		p=p.Next
		count++
	}
	return p
}

//获取数组长度
func GetLength(l *List.ListNode) int {
	var length int
	p:=l
	for p!=nil {
		length++
		p=p.Next
	}
	return length
}
