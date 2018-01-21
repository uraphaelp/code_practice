package ransomenote2

func Canconstruct(ransome string, magazine string) bool {
	map1:=make(map[byte]int)
	for i:=0; i<len(magazine); i++ {
		map1[magazine[i]]++
	}
	for i:=0; i<len(ransome); i++ {
		map1[ransome[i]]--
		if map1[ransome[i]]<0 {
			return false
		}
	}
	return true
}
