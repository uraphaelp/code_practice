## 字符串换序输出

## 题意分析
`PAYPALISHIRING` ,设置 `numrows=3` ,输出为：
```
P   A   H   N
A P L S I I G
Y   I   R
```
即：`PAHNAPLSIIGYIR`

### **zigzag概念分析**
```
/*n=numRows
Δ=2n-2    1                           2n-1                         4n-3
Δ=        2                     2n-2  2n                    4n-4   4n-2
Δ=        3               2n-3        2n+1              4n-5       .
Δ=        .           .               .               .            .
Δ=        .       n+2                 .           3n               .
Δ=        n-1 n+1                     3n-3    3n-1                 5n-5
Δ=2n-2    n                           3n-2                         5n-4
*/
```
## 算法分析
单列 **从上至下** 导入数据，显然需要在 `s[i](i=numrows-1)` 时 **转换** 导入数据的方向，直至回到 `第0行` 再次 **换向**
