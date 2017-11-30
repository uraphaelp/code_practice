func twoSum(numbers []int, target int) []int {
    left:=0
    right:=len(numbers)-1
    for left<right {
        if numbers[left]+numbers[right]==target {
            return []int{left+1, right+1}
        //以下步骤保证每个数都能遍历到
        } else if numbers[left]+numbers[right]<target {
            left++
        } else {
            right--
        }
    }
    return []int{}
}
