//direct and simple but time-wasting solution
func mySqrt(x int) int {
    if x==0 {
        return 0
    }    
    for i:=1; ; i++ {
        tmp:=x/i
        if tmp>i {
            continue
        } else if tmp==i {
            return i
        } else if tmp<i && i*i<x {
            return i
        } else {
            return i-1
        }
    }
}

//using binary search
func mySqrtII(x int) int {
    if x <= 1 {
        return x
    }
    low := 1
    high := x
    for low < high {
        mid := (low + high) / 2
        sq := mid*mid
        if sq > x {
            high = mid
        } else if sq < x {
            low = mid + 1
        } else {
            return mid
        }
    } 
    return high-1
}
