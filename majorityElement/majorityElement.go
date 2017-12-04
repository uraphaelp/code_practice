type Slice []int

func (arr Slice) Len() int {
    return len(arr)
}

func (arr Slice) Swap(i, j int) {
    arr[i], arr[j]=arr[j], arr[i]
}

func (arr Slice) Less(i, j int) bool {
    return arr[i]<arr[j]
}
/*以上为了使用sort包*/

func majorityElement(nums Slice) int {
    l:=len(nums)
    if !sort.IsSorted(nums) {
        sort.Sort(nums)
    }
    return nums[l/2]
}
