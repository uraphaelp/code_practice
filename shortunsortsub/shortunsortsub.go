func Findunsortsub(nums []int) int {
	var (
		left, right, min, max int
		tmpsub []int
	)
  //粗糙定位left, right
	for i:=1; i<len(nums); i++ {
		if nums[i]<nums[i-1] {
			left=i-1
			break
		}
	}
	for j:=len(nums)-1; j>0; j--{
		if nums[j]<nums[j-1] {
			right=j
			break
		}
	}
  //找出必定需要重排的tmpsub
	tmpsub=nums[left:right+1]
	min, max=Findminandmax(tmpsub)
  //确定min, max of tmpsub在两侧有序序列的位置
	for i:=0; i<len(nums[0:left+1]); i++ {
		if min<nums[i] {
			left=i
		}
	}
	for j:=len(nums)-1; j>right-1; j-- {
		if max>nums[j] {
			right=j
		}
	}
	if right==left {
		return 0
	}
	return right-left+1
}

//辅助函数：找出数组中最小、最大元素
func Findminandmax(nums []int) (int, int) {
	min:=nums[0]
	max:=nums[0]
	if len(nums)==0 {
		return min, max
	}
	for i:=1; i<len(nums); i++ {
		if nums[i]<min {
			min=nums[i]
		}
		if nums[i]>max {
			max=nums[i]
		}
	}
	return min, max
}
