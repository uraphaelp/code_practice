func Finddisappear(nums []int) []int {
	var (
		res []int
		val int
	)
	for i:=0; i<len(nums); i++ {
		if nums[i]<0 {
			val=-nums[i]-1
		} else {
			val=nums[i]-1
		}
		if nums[val]>0 {
			nums[val]*=-1
		}
	}
	for i:=0; i<len(nums); i++ {
		if nums[i]>0 {
			res=append(res, i+1)
		}
	}
	return res
}
