// 对输入字符串的非空格部分计数："hello ,my name is Raphael" (count=5)
func Countsegments(s string) int {
    count:=0
	  for i:=0; i<len(s); i++ {
		    if s[i]!=32 && (i==0 || s[i-1]==32) {
			      count++
		    }
	  }
	  return count
}
