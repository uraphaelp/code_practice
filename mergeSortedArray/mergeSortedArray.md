# mergeSortedArray的“bug”
- ###  **当测试程序为：**
```
arr1 := []int{1, 3, 5, 7, 9}
arr2 := []int{2, 4, 6, 8, 10}
merge(arr1, 5, arr2, 5)
fmt.Println(arr1)
```
会出现 `runtime error: slice bounds out of range` 原因是当切片组长度为5，不可添加多于5个元素
- ###  **若将测试程序改为：**
```
arr1 := []int{1, 3, 5, 7, 9}
arr2 := []int{2, 4, 6, 8, 10}
merge(arr1, 2, arr2, 3)
fmt.Println(arr1)
```
Output: [1, 2, 3, 4, 6]

原因是当前merge后的元素个数为5，arr1可以容纳。本次merge操作相当于将 `[]int{1, 3}, []int{2, 4, 6}` 拼接

**总之，arr1原长度需大于m+n，方可完成正确的merge**
