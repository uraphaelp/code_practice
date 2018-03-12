## 查找旋转数组中的target

## 算法分析
1 要求时间复杂度为 `O(log n)`

2 考虑先二分查找定出orotate element位置

3 然后对比 `target & nums[0]` 左边查找，右边查找

4 第3步容易出现边界问题：如数组未旋转，则定出的 `rotate eleindex=0`， 对于如 `Search(nums=[1,3], 3)` 

```
由于target==3>nums[0]==1, 所以根据第3步，在左边，即nums[0-0]查找，则返回-1，与期待的1不符合
所以考虑应该找到 rotate eleindex 后不拆分成两部分
```
