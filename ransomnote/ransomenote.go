//magazine中是否含有ransome字符串
package ransomnote

func Canconstruct(ransome string, magazine string) bool {
	//round标记是否需要重新(开始)计数
	var j, round, count int
	if len(ransome)==0 {
		return true
	}
	for i:=0; i<len(magazine); i++ {
		if magazine[i]==ransome[j] {
			j++
			round=0
			count++
		} else if magazine[i]!=ransome[j] && round==0 {
			i--
			j=0
			round=1
			count=0
		} else {
			j=0
			round=0
		}
		if count==len(ransome) {
			return true
		}
	}
	return false
}
