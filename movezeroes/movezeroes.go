//将数组的所有0移到末尾
func moveZeroes(nums []int)  {
    var i, j, count, l int
	l = len(nums)
	count = l
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
			//删除这个元素
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	//把所有0放到末尾
	for j = 0; j < l-count; j++ {
		nums = append(nums, 0)
	}
}