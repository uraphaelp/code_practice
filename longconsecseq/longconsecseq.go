//找出乱序数组中最长的连续序列
func Longestconsecseq(nums []int) int {
	var (
		res, left, right, x, sum int
	)
	m:=make(map[int]int)
	for _,x=range nums {
		if m[x]==0 {
			if m[x-1]==0 {
				left=0
			} else {
				left=m[x-1]
			}
			if m[x+1]==0 {
				right=0
			} else {
				right=m[x+1]
			}
			sum=left+right+1
			m[x]=sum
			if sum>res {
				res=sum
			}
			//core!!!core!!!
			m[x-left]=sum
			m[x+right]=sum
		} else {
			continue
		}
	}
	return res
}
