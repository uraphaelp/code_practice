func Swappairs(head *List.ListNode) *List.ListNode {
	if head==nil || head.Next==nil {
		return head
	}
	cur:=head
	newhead:=head.Next
	for cur!=nil && cur.Next!=nil {
		tmp:=cur
		cur=cur.Next
		tmp.Next=cur.Next
		cur.Next=tmp
		cur=tmp.Next
		if cur!=nil && cur.Next!=nil {
			tmp.Next = cur.Next
		}
	}
	return newhead
}
