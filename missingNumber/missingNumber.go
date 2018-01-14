func missingNumber(nums []int) int {
	var practicalsum, realsum, l int
	l = len(nums)
	for i := 0; i < l; i++ {
		practicalsum += nums[i]
	}
	realsum = (1 + l) * l / 2
	return (realsum - practicalsum)
}
