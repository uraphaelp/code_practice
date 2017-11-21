//simple and direct solution
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
