import "sort"

func Maxofthree(nums []int) int {
	l:=len(nums)
	sort.Ints(nums)
	return Findmax(nums[0]*nums[1]*nums[l-1], nums[l-3]*nums[l-2]*nums[l-1])
}

func Findmax(a, b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}
