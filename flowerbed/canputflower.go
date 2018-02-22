var (
		i, count int
		l=len(flowerbed)
	)
	for i=0; i<l; i++ {
    //注意这个满足的条件
		if flowerbed[i]==0 && (i==0 || flowerbed[i-1]==0) && (i==l-1 || flowerbed[i+1]==0) {
      //表示该位置种下
			flowerbed[i]=1
			count++
		}
	}
	if count>=n {
		return true
	} else {
		return false
	}
