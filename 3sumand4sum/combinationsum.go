//找到数组中任意大小组合之和为target的组合
func Combinationsum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var (
		res=make([][]int, 0, 0)
		tmp=make([]int, 0, 0)
	)
	dfs(candidates, target, 0, &tmp ,&res)
	return res
}

func dfs(nums []int, target, start int, onedim *[]int, twodim *[][]int) {
	if target==0 {
		tmp1:=make([]int, len(*onedim), len(*onedim))
		copy(tmp1, *onedim)
		*twodim=append(*twodim, tmp1)
		return
	}
	if target<0 {
		return
	}
	for i:=start; i<len(nums); i++ {
		if i>start && nums[i]==nums[i-1] {
			continue
		}
		*onedim=append(*onedim, nums[i])
		dfs(nums, target-nums[i], i+1, onedim, twodim)
		*onedim=append((*onedim)[:len(*onedim)-1], (*onedim)[len(*onedim):]...)
	}
}
