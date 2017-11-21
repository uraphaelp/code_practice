//simple and direct solution: from the bottom
/*从数组后部开始，依次降序插入元素。若从头开始，会覆盖原数组元素*/

func merge(nums1 []int, m int, nums2 []int, n int)  { 
    for n>0 {
        if m==0 || nums2[n-1]>nums1[m-1] {
            nums1[m+n-1]=nums2[n-1]
            n--
        } else {
            nums1[m+n-1]=nums1[m-1]
            m--
        }
    }    
}

//different idea：from the head
func mergeII(nums1 []int, m int, nums2 []int, n int) {
    copy(nums1[m:m+n], nums2[:n]) //把nums2元素放入nums1后部
    left, mid, right := 0, m-1, m
    for left <= mid && right <= m+n-1 {
        if nums1[left] <= nums1[right] { //if left不大于right，元素位置不变，left+1           
            left++
        } else {  //if right<left, right替换到当前left位置，并且从当前left位置开始，所有数组元素后移一位
            tmp := nums1[right]
            copy(nums1[left+1:right+1], nums1[left:right])
            nums1[left] = tmp
            left++
            mid++ //记得mid++，因为left元素已经后移了一位
            right++
        }
    }
}
