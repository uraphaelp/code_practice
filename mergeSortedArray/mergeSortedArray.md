# 合并两有序数组，同样按照升序/降序排列(这里是升序)
## mergeSortedArray解析

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.

注意：实际上题意并不要求m, n分别等于两切片元素的个数。本题要求：
- m,n 应分别小于各自切片中的实际元素个数
- arr1的容量应>m+n

## mergeSortedArray的“bug”
- ###  **当测试程序为：**
```
arr1 := []int{1, 3, 5, 7, 9}
arr2 := []int{2, 4, 6, 8, 10}
merge(arr1, 5, arr2, 5)
fmt.Println(arr1)
```
会出现 `runtime error: slice bounds out of range` 原因是当切片组长度为5，不可添加多于5个元素

若想完成1-10的升序排列，则应`arr1=make([]int, 10, 10)` ，同时`arr1[0:5]=T  T:=[]int{1, 2, 3, 4, 5}`

此时`arr1=[1, 2, 3, 4, 5, 0, 0, 0, 0, 0]`，merge过程同下
- ###  **若将测试程序改为：**
```
arr1 := []int{1, 3, 5, 7, 9}
arr2 := []int{2, 4, 6, 8, 10}
merge(arr1, 2, arr2, 3)
fmt.Println(arr1)
```
Output: [1, 2, 3, 4, 6]

原因是当前merge后的元素个数为5，arr1可以容纳。本次merge操作相当于将 `[]int{1, 3}, []int{2, 4, 6}` 拼接

**总之，arr1原长度需大于m+n，方可完成正确的merge**
