## 数组原本含有1-n的元素. 寻找数组中丢失的元素：某些元素可能出现两次，而有些元素可能丢失，找出丢失的元素

### 不额外使用空间，利用O(n)时间复杂度完成
以数组[4, 3, 2, 7, 8, 2, 3, 1]为例：

实际上以数组中每个元素值val作为下标，该元素nums[val]-1都存在于数组中

则在第一轮循环中将nums[val]-1置负，第二轮循环查找，若元素值为负，则说明却少的数为i-1