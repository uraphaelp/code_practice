func Checkperfect(num int) bool {
	var sum int
	if num==0 {
		return false
	}
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
