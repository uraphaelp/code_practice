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

func majorityElementII(nums []int) int {
    m := make(map[int]int)   //using map
    l := len(nums)
    for _, n := range nums { //n for nums[i] value
        m[n]++               //m[n] for times of n apperence
        if m[n] > l/2 {
            return n
        }
    }
    return 0
}

func majorityElementIII(nums []int) int {
    var count, majority, i int
    l:=len(nums)
    majority=nums[0]
    count=1
    for i=1; i<l; i++ {
        if count==0 {
            majority=nums[i]
            count++
        } else if nums[i]==count {
            count++
        } else {
            count--
        }
    }
    return majority
}
