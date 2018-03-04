//LCIS=Longest Continous Increase Subarray
func FindlengthofLCIS(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	var (
		count, res=1, 1
	)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i]<nums[i+1] {
			count++
			if res<count {
				res=count
			}
		} else {
			count=1
		}
	}
	return res
}
