//将left---right范围内所有：可以除尽自身组成数字的数返回
//如left=1, right=22，则返回：[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]

func Selfdividingnums(left, right int) []int {
	var result []int
	for i:=left; i<=right; i++ {
		if Isselfdiving(i) {
			result=append(result, i)
		}
	}
	return result
}

func Isselfdiving(num int) bool {
	result:=num
	for result>0 {
		if result%10==0 || num%(result%10)!=0 {
			return false
		}
		result/=10
	}
	return true
}
