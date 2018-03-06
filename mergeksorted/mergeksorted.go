/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var (
		index, value []int
		min, seq int
	)
	res:=new(ListNode)
	res.Val=0
	res.Next=nil
	for i,v:=range lists {
		if v!=nil {
			index=append(index, i)
			value=append(value, v.Val)
		}
	}
	for len(value)!=0 {
		min, seq = Findmin(value)
		for k, j := range index {
			if lists[j].Val == min {
				node := new(ListNode)
                value=append(value[:seq], value[seq+1:]...)
				node.Val = min
				node.Next = nil
				Findlast(res).Next = node
				lists[j] = lists[j].Next
				if lists[j]!=nil {
					value = append(value, lists[j].Val)
                } else {
                    index=append(index[:k], index[k+1:]...)
                }
				break
			}
		}
	}
	return res.Next
}

func Findmin(nums []int) (int, int) {
    var index int
	min:=nums[0]
	for i:=0; i<len(nums); i++ {
		if nums[i]<min {
			min=nums[i]
            index=i
		}
	}
	return min, index
}

func Findlast(head *ListNode) *ListNode {
	p:=head
	for p.Next!=nil {
		p=p.Next
	}
	return p
}
