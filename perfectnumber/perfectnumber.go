//求num的所有因子的和?=num
func Checkperfect(num int) bool {
	var sum int
	if num==0 {
		return false
	}
	//注意这个阈值条件的设置
	for i:=1; i*i<=num; i++ {
		if num%i==0 {
			sum+=i
			if i*i!=num {
				sum+=num/i
			}
		}
	}
	return sum-num==num
}
