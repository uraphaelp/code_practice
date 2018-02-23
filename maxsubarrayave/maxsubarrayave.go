func Findmaxaverage(nums []int, k int) float64 {
	var (
		sum float64
		res float64
	)
  //注意滑动窗口的设置
	for i:=0; i<k; i++ {
		sum+=float64(nums[i])
	}
	res=sum
	for i:=k; i<len(nums); i++ {
		sum=sum+float64(nums[i]-nums[i-k])
		res=maxoftwo(sum, res)
	}
	return res/float64(k)
}

func maxoftwo(a, b float64) float64 {
	if a>=b {
		return a
	}
	return b
}
