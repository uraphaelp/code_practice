func isUgly(num int) bool {
    if num<1 {
        return false
    }
    for i:=2; i<6 && num>0; i++ { //质因子为2、3、5(4无关紧要，相当于除了两个2)
        for num%i==0 {
            num/=i  //不断消除2,3,5这几个质因子，剩下其他的因子：查看是否有>=7的，如果有则必然除不尽这几个数
        }
    }
    return num==1
}

//下述方法拥有与上述一致的思路
func isUgly(num int) bool {
    if num == 1 {
		    return true
	  }
    if num <= 0 {
        return false
    }
	  for num > 1 {
		    if num%2 == 0 {
			      num = num / 2
		    } else if num % 3 == 0 {
			      num = num / 3
		    } else if num % 5  == 0{
			      num = num / 5
		    } else {
  		      return false
        }
	  }
	  return true    
}
