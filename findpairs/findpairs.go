//给出nums, 目标k, 寻找数组中有多少组任意两个数之差为 k
//[3, 1, 4, 1, 5], k=2: 2组：[3,1], [3,5] 重复的不计入
//[1, 2, 3, 4, 5], k=1: 4组
//[1, 3, 1, 5, 4], k=0: 1组

func Findpairs(nums []int, k int) int {
	if k<0 {
		return 0
	}
	var res int
	m:=make(map[int]int)
	for _, v:=range nums {
		m[v]++
	}
	for _,v:=range nums {
      //注意k=0的情况!!!
			if k==0 {
			if m[v]>1 {
				res++
				m[v]=0
			}
		} else if m[v+k]!=0 {
			res++
			m[v+k]=0
		}
	}
	return res
}
