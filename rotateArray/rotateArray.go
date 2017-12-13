func rotate(nums []int, k int)  {
    var tmp []int
    l:=len(nums)
    if k>=l {
        k=k%l
    }
    if k==0 {
        return
    } else {
        for i:=0; i<len(nums[0:l-k]); i++ {
            tmp=append(tmp, nums[i])
        }
        nums=append(nums[:0], nums[l-k:]...)
        nums=append(nums, tmp...)
    }
}
