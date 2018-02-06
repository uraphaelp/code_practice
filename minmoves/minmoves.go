func Minmoves(nums []int) int {
if len(nums)==0 {
		return 0
	}
	var (
		sum int
		min=nums[0]
	)
	for i:=0; i<len(nums); i++ {
		sum+=nums[i]
		if nums[i]<min {
			min=nums[i]
		}
	}
	return sum-min*(len(nums)-1)
}
