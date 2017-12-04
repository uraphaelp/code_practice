func titleToNumber(s string) int {
    var i, l, result int
    l=len(s)
    result=0
    for i=0; i<l; i++ {
        /*注意：不要从末尾开始乘26的低次幂，从前面开始乘，每次都乘26，相当于高位就能在每次循环时多乘一个26*/
        result=result*26+int(s[i]-'A')+1
    }
    return result
}
