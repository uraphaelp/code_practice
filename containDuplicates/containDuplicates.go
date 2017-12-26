func containsNearbyDuplicate(nums []int, k int) bool {
    m:=make(map[int]int)
    var left, right int
    if k<=0 {
        return false
    }
    if k>len(nums)-1 {
        k=len(nums)-1
    }
    for i:=0; i<len(nums); i++ {
        v, _:=m[nums[i]]
        //v存在与否：
        if v==0 {
            m[nums[i]]++
            right++
        } else {
            return true
        }
        //若超过容限，删除前面的元素
        if right-left>k {
            m[nums[left]]=0
            left++
        }
    }
    return false
}
