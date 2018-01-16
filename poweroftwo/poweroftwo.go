//判断是否为2的幂次
import (
    "fmt"
    "math"
)

//使用二进制的特点：若n为2的幂次，则形式必为：100000...形式；n-1则为11111...形式
func isPowerOfTwo1(n int) bool {
    if n <= 0 {
        
        return false
    }
    if n == 1 {
        return true
    }
	//n & n-1 必为0(没有任何一位相同)
    return ( n & (n-1)) == 0
}

//常规方法
func isPowerOfTwo1(n int) bool {
	if n<=0 {
        return false
    }    
    for n>2 {
        if n%2!=0 {
            return false
        }
        n/=2
    }
    return true
}