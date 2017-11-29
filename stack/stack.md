# 栈的实现
## 一种主函数赋值方式
```
var stack Stack
var a, b, c, d, e, f ListNode
a.Val=1
a.Next=&b
......
f.Val=1
f.Next=nil
stack.Head=&a
stack.Bottom=f.Next
```
