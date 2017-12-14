func isHappy(n int) bool {
    //n from [2,6] is not happy
    for n>6 {
        next:=0
        for n!=0 {
            next+=(n%10)*(n%10)
            n/=10
        }
        n=next
    }
    return n==1
}
