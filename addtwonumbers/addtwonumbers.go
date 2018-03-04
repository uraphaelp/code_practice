// 两个多位数以链表形式存储(eg. 3-4-2:243)，返回两个数相加的结果

import "github.com/uraphaelp/List"

func Addtwonumbers(l1 *List.ListNode, l2 List.ListNode) *List.ListNode {
	res:=new(List.ListNode)
	var (
		carry, pos int
	)
	if l1!=nil || l2!=nil {
		if l1==nil {
			l1.Val=0
		}
		if l2==nil {
			l2.Val=0
		}
		pos=l1.Val+l2.Val
		if pos>=10 {
			carry=pos/10
			pos=pos%10
		} else {
			carry=0
		}
	}
	l1=l1.Next
	l2=l2.Next
	res.Val=pos
	res.Next=nil
	for l1!=nil || l2!=nil || carry!=0 {
		if l1==nil {
			l1.Val=0
		}
		if l2==nil {
			l2.Val=0
		}
		pos=l1.Val+l2.Val+carry
		if pos>=10 {
			carry=pos/10
			pos=pos%10
		} else {
			carry=0
		}
		l1=l1.Next
		l2=l2.Next
		InserttoTail(res, pos)
	}
	return res
}

//从尾部插入
func InserttoTail(l *List.ListNode, x int)  {
	p:=FindTail(l)
	s:=new(List.ListNode)
	s.Val=x
	s.Next=nil
	p.Next=s
}

func FindTail(l *List.ListNode) *List.ListNode {
	p:=l
	for p.Next!=nil {
		p=p.Next
	}
	return p
}
