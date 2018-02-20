//二进制数组：找出之中 1 的最大连续次数，返回连续次数

func Findmaxconsecutiveones(nums []int) int {
var (
		count, max int
    //flag作为标记，若扫描到0，flag归0，否则置1
		flag=1
	)
	for i:=0; i<len(nums); i++ {
		if nums[i]==1 && flag==1 {
			count++
		} else if nums[i]==1 && flag==0 {
			count++
      //flag需要重新置1
			flag=1
		} else {
			flag=0
			count=0
		}
		if count>max {
			max=count
		}
	}
	return max
}
