func Reversekgroup(head *ListNode, k int) *ListNode {
	n:= 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	dmy := new(ListNode)
	dmy.Val=0
	dmy.Next = head
	s := dmy 
	e := dmy.Next //s: start, e: end
	for i := n; i >= k; i -= k {
		for  j := 1; j < k; j++ { // reverse group
			Next := e.Next
			e.Next = Next.Next
			Next.Next = s.Next
			s.Next = Next
		}
		s = e
		e = s.Next
	}
	return dmy.Next
}
