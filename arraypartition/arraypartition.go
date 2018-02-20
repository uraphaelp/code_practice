//给定一个2n个整数的数组，寻找一个方法使得数组中的数两两组合(a1, b1)...(an bn)，最终min(a1,b1)+...+min(an, bn)最大
//算法：排序后相隔的数相加即可得到最大

func Arrayparsum(nums []int) int {
	var res int
	sort.Ints(nums)
	for i:=0; i<len(nums); i+=2 {
		res+=nums[i]
	}
	return res
}
