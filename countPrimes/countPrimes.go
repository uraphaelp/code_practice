func countPrimes(n int) int {
    var primes []bool
    //全部设定为true
    primes=[]bool{true, true}
    var count int
    //全部以true存进切片
    for i:=2; i<n; i++ {
        primes=append(primes, true)
    }
    for i:=2; i*i<n; i++ {
        if !primes[i] {
            continue
        }
        //为了防止重复判定，从i^2开始每间隔i个，这些都不是质数
        for j:=i*i; j<n; j+=i {
            primes[j]=false
        }
    }
    for i:=2; i<n; i++ {
        if primes[i] {
            count++
        }
    }
    return count
}
