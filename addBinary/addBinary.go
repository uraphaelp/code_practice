func addBinary(a string, b string) string {
    i:=len(a)-1
    j:=len(b)-1
    c:=0
    s:=""
    for i>=0 || j>=0 || c==1 {
    /*两个数未取完，或进位为1*/
        if i>=0 {
            c+=int(a[i]-'0')
            i--
        }
        if j>=0 {
            c+=int(b[j]-'0')
            j--
        }
        s=strconv.Itoa(c%2)+s
         /*每一位相加完都写入字符串，避免进位和当前数位的混乱*/
        c/=2
        /*设置进位：每一位相加只有两种结果：0、1、2(10)。其中前两种进位为0，后一种进位为1*/
        /*不保留当前数位，算完一位写入一次，丢弃当前数位，只保留进位*/
    }
    return s
}
