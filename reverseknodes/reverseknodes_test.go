package reversenodes_k

import (
	"testing"
	"fmt"
)
func TestReverse(t *testing.T) {
	a:=new(ListNode)
	b:=new(ListNode)
	c:=new(ListNode)
	d:=new(ListNode)
	e:=new(ListNode)
	a.Val=1
	a.Next=b
	b.Val=2
	b.Next=c
	c.Val=3
	c.Next=d
	d.Val=4
	d.Next=e
	e.Val=5
	e.Next=nil
	fmt.Println(Traverse(Reverse(a)))
}

func Traverse(head *ListNode) []int {
	var res []int
	p:=head
	for p!=nil {
		res=append(res, p.Val)
		p=p.Next
	}
	return res
}
