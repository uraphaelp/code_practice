# 最大利润
## 题意解析
- Example
```
Input: [7, 1, 5, 3, 6, 4]
Output: 5
note: max=6-1=5
```
## 算法描述
找最大利润：
- **遍历时，后面的一个元素减去前面的一个元素，结果>0, 当做盈利，比较之前得出的盈利：若大则更新盈利值，若小则不考虑**
- **结果<0, 找到一个更小的对比值，更新对比值**

![算法描述](https://github.com/uraphaelp/code_practice/blob/master/maxProfit/maxprofit.jpg)
