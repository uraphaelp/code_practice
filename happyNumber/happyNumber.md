# 是否是happy number
## 题意分析
```
19是happy number，因为：
1^2+9^2=82
8^2+2^2=68
6^2+8^2=100
1^2+0^2+0^2=1
```
## 算法解析
- **[2,6]均不是happy number；所以一旦某一步出现这5个数之中的一个，即判断它不是happy number**
```
for n>6 {
    for n!=0 {
        next+=(n%10) * (n%10) 
        n/10
     }
     n=next
}
return n==1 //在n<=6中，只有1为happy
