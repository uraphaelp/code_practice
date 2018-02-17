func judgeSquareSum(c int) bool {
    if c==0 {
        return true
    }
    var (
		i int
		b int
	)
	for i=0; i*i<c; i++ {
		b=c-(i*i)
		if Binarysearch(0,b,b) {
			return true
		}
	}
	return false
}

//n是目标数
func Binarysearch(left, right, n int) bool {
	var mid int
	if left>right {
		return false
	}
	mid=left+(right-left)/2
	if mid*mid==n {
		return true
	} else if mid*mid>n {
		return Binarysearch(left, mid-1, n)
	} else {
		return Binarysearch(mid+1, right, n)
	}
}
