func Checkpossibility(nums []int) bool {
	var (
		index=-1
	)
  //找出逆序的index；同时若有相邻的超过2个以上的逆序对，则false
	for i:=0; i<len(nums)-1; i++ {
		if nums[i]>nums[i+1] {
			if index!=-1 {
				return false
			}
			index=i
		}
	}
  //这几种true: -1(没有逆序对), 0(只有一个逆序对), (1个逆序对，出现在最后两个元素), (逆序位置左右两侧元素顺序，则只需改变逆序位置)
  //(易忽略：nums[index]<nums[index+1] && nums[index]<nums[index+2])
	if index==-1 || index==0 || index==len(nums)-2 || nums[index-1]<=nums[index+1] || nums[index]<=nums[index+2] {
		return true
	}
	return false
}
